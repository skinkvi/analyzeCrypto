package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/skinkvi/analyzeCrypto/internal/logger"
)

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", h.healthCheck)
}

func (h *Handler) healthCheck(c *gin.Context) {
	logger.Logger.Info("Запрос на /health")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
