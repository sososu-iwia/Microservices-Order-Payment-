package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"order-service/internal/domain"
	repo "order-service/internal/repository/postgres"
	transport "order-service/internal/transport/http"
	"order-service/internal/usecase"
)

type Config struct{ PaymentBaseURL string }
type RouterDeps struct {
	DB     *sql.DB
	Config Config
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now().UTC() }

type uuidGenerator struct{}

func (uuidGenerator) NewID() string { return uuid.NewString() }

type orderUsecaseAdapter struct{ uc *usecase.OrderUsecase }

func (a *orderUsecaseAdapter) CreateOrder(ctx *gin.Context, customerID, itemName string, amount int64, idempotencyKey string) (*domain.Order, int, error) {
	return a.uc.CreateOrder(ctx.Request.Context(), customerID, itemName, amount, idempotencyKey)
}
func (a *orderUsecaseAdapter) GetOrder(ctx *gin.Context, id string) (*domain.Order, error) {
	return a.uc.GetOrder(ctx.Request.Context(), id)
}
func (a *orderUsecaseAdapter) CancelOrder(ctx *gin.Context, id string) (*domain.Order, error) {
	return a.uc.CancelOrder(ctx.Request.Context(), id)
}

func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.Default()
	orderRepo := repo.NewOrderRepository(deps.DB)
	paymentClient := NewPaymentHTTPClient(deps.Config.PaymentBaseURL, &http.Client{Timeout: 2 * time.Second})
	uc := usecase.NewOrderUsecase(orderRepo, paymentClient, realClock{}, uuidGenerator{})
	handler := transport.NewHandler(&orderUsecaseAdapter{uc: uc})
	handler.RegisterRoutes(r)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	return r
}
