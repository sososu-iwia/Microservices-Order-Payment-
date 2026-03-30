package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"payment-service/internal/domain"
	repo "payment-service/internal/repository/postgres"
	transport "payment-service/internal/transport/http"
	"payment-service/internal/usecase"
)

type RouterDeps struct{ DB *sql.DB }

type uuidGenerator struct{}

func (uuidGenerator) NewID() string { return uuid.NewString() }

type paymentUsecaseAdapter struct{ uc *usecase.PaymentUsecase }

func (a *paymentUsecaseAdapter) CreatePayment(ctx *gin.Context, orderID string, amount int64) (*domain.Payment, error) {
	return a.uc.CreatePayment(ctx.Request.Context(), orderID, amount)
}
func (a *paymentUsecaseAdapter) GetByOrderID(ctx *gin.Context, orderID string) (*domain.Payment, error) {
	return a.uc.GetByOrderID(ctx.Request.Context(), orderID)
}

func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.Default()
	paymentRepo := repo.NewPaymentRepository(deps.DB)
	uc := usecase.NewPaymentUsecase(paymentRepo, uuidGenerator{})
	handler := transport.NewHandler(&paymentUsecaseAdapter{uc: uc})
	handler.RegisterRoutes(r)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	return r
}
