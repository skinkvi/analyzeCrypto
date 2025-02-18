package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/skinkvi/analyzeCrypto/internal/config"
	"github.com/skinkvi/analyzeCrypto/internal/logger"
	"go.uber.org/zap"
)

// Ф-ия инициализации бд с проверкой (ping)
func InitDB(cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Logger.Error("Ошибка парсинга конфигурации подключения", zap.Error(err))
		return nil, errors.Wrap(err, "ошибка парсинга конфигурации подключения")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Logger.Error("Ошибка подлкючения к базе данных", zap.Error(err))
		return nil, errors.Wrap(err, "ошибка подлкючения к базе данных")
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Logger.Error("Ошибка проверки подключения к базе данных", zap.Error(err))
		return nil, errors.Wrap(err, "ошибка проверки подключения к базе данных")
	}

	logger.Logger.Info("База данных успешно подключилась")

	return pool, nil
}
