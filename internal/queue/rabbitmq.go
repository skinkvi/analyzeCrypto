package queue

import (
	"context"

	"github.com/skinkvi/analyzeCrypto/internal/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Logger.Error("Ошибка подключение к RabbitMQ", zap.Error(err))
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Logger.Error("Ошибка создания канала в RabbitMQ", zap.Error(err))
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Publish(ctx context.Context, queueName string, body []byte) error {
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
		return err
	}

	return nil
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
