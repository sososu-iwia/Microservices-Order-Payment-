package postgres

import (
	"context"
	"database/sql"
	"payment-service/internal/domain"
)

type PaymentRepository struct{ db *sql.DB }

func NewPaymentRepository(db *sql.DB) *PaymentRepository { return &PaymentRepository{db: db} }

func (r *PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO payments (id, order_id, transaction_id, amount, status)
        VALUES ($1, $2, $3, $4, $5)
    `, payment.ID, payment.OrderID, payment.TransactionID, payment.Amount, payment.Status)
	return err
}

func (r *PaymentRepository) GetByOrderID(ctx context.Context, orderID string) (*domain.Payment, error) {
	row := r.db.QueryRowContext(ctx, `
        SELECT id, order_id, transaction_id, amount, status
        FROM payments
        WHERE order_id = $1
        ORDER BY created_at DESC
        LIMIT 1
    `, orderID)
	var payment domain.Payment
	if err := row.Scan(&payment.ID, &payment.OrderID, &payment.TransactionID, &payment.Amount, &payment.Status); err != nil {
		return nil, err
	}
	return &payment, nil
}
