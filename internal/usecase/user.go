package usecase

import (
	"github.com/skinkvi/analyzeCrypto/internal/queue"
	"github.com/skinkvi/analyzeCrypto/internal/repository/postgres"
)

type UserUsecase struct {
	repo     postgres.UserRepository
	rabbitMQ *queue.RabbitMQ
}

func NewUserUsecase(repo postgres.UserRepository, rabbitMQ *queue.RabbitMQ) *UserUsecase {
	return &UserUsecase{
		repo:     repo,
		rabbitMQ: rabbitMQ,
	}
}
