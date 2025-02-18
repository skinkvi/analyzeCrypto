package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/skinkvi/analyzeCrypto/internal/config"
	"github.com/skinkvi/analyzeCrypto/internal/errors"
	"github.com/skinkvi/analyzeCrypto/internal/logger"
	"go.uber.org/zap"
)

const pkg = "db"

// Ф-ия инициализации бд с проверкой (ping)
func InitDB(cfg *config.Config) (*pgxpool.Pool, error) {
	const method = "InitDB"

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Logger.Error("Ошибка парсинга конфигурации подключения", zap.Error(err))
		return nil, errors.Wrap(err, pkg, method, "Ошибка парсинга конфигурации подключения")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Logger.Error("Ошибка подлкючения к базе данных", zap.Error(err))
		return nil, errors.Wrap(err, pkg, method, "Ошибка подключения к базе данных")
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Logger.Error("Ошибка проверки подключения к базе данных", zap.Error(err))
		return nil, errors.Wrap(err, pkg, method, "ошибка проверки подключения к базе данных")
	}

	logger.Logger.Info("База данных успешно подключилась")

	return pool, nil
}
