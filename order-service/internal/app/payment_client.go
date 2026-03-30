package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"order-service/internal/usecase"
)

type PaymentHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewPaymentHTTPClient(baseURL string, client *http.Client) *PaymentHTTPClient {
	return &PaymentHTTPClient{baseURL: baseURL, client: client}
}

type paymentRequest struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}
type paymentResponse struct {
	ID            string `json:"id"`
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"`
	Amount        int64  `json:"amount"`
	Status        string `json:"status"`
}

func (c *PaymentHTTPClient) CreatePayment(ctx context.Context, orderID string, amount int64) (*usecase.PaymentResult, error) {
	body, _ := json.Marshal(paymentRequest{OrderID: orderID, Amount: amount})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/payments", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result paymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &usecase.PaymentResult{Status: result.Status, TransactionID: result.TransactionID}, nil
}
