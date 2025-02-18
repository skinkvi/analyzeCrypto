package main

import (
	"context"
	"embed"
	"io/fs"

	"github.com/jackc/tern/v2/migrate"
	"github.com/skinkvi/analyzeCrypto/internal/config"
	"github.com/skinkvi/analyzeCrypto/internal/db"
	"github.com/skinkvi/analyzeCrypto/internal/logger"
	"go.uber.org/zap"
)

var migrationsFS embed.FS

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()

	cfg := config.LoadConfig()

	pool, err := db.InitDB(cfg)
	if err != nil {
		logger.Logger.Fatal("Ошибка инициализации базы данных в мейне", zap.Error(err))
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		logger.Logger.Fatal("Ошибка получения подключения в мейне", zap.Error(err))
	}
	defer conn.Release()

	migrator, err := migrate.NewMigrator(context.Background(), conn.Conn(), "schema_migrations")
	if err != nil {
		logger.Logger.Fatal("Ошибка создания мигратора", zap.Error(err))
	}

	// Используем embed.FS для загрузки миграций
	migrationsDir, err := fs.Sub(migrationsFS, "migrations")
	if err != nil {
		logger.Logger.Fatal("Ошибка загрузки миграций", zap.Error(err))
	}

	err = migrator.LoadMigrations(migrationsDir)
	if err != nil {
		logger.Logger.Fatal("Ошибка загрузки миграций", zap.Error(err))
	}

	err = migrator.Migrate(context.Background())
	if err != nil {
		logger.Logger.Fatal("Ошибка применения миграций", zap.Error(err))
	}

	logger.Logger.Info("Миграции успешно применены")
}
