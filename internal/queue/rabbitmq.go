package queue

import (
	"context"

	"github.com/skinkvi/analyzeCrypto/internal/errors"
	"github.com/skinkvi/analyzeCrypto/internal/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

const pkg = "queue"

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	const method = "NewRabbitMQ"

	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Logger.Error("Ошибка подключение к RabbitMQ", zap.Error(err))
		return nil, errors.Wrap(err, pkg, method, "Ошибка подключения к RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Logger.Error("Ошибка создания канала в RabbitMQ", zap.Error(err))
		return nil, errors.Wrap(err, pkg, method, "Ошибка создания канала в RabbitMQ")
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Publish(ctx context.Context, queueName string, body []byte) error {
	const method = "Publish"

	err := r.ch.Publish(
		"",        // обмен
		queueName, // ключ маршрутизации
		false,     // обязательное?
		false,     // немедленное?
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		logger.Logger.Error("Ошибка публикации сообщения в очередь RabbitMQ", zap.Error(err))
		return errors.Wrap(err, pkg, method, "Ошибка публикации сообщения в очередь RabbitMQ")
	}

	return nil
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
