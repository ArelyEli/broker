package schemas

type CreateTransactionRequest struct {
	Amount     float64 `json:"amount" binding:"required" required:"$field is required"`
	MerchantID uint    `json:"merchantID" binding:"required" required:"$field is required"`
}

type TransactionResponse struct {
	ID         uint    `json:"transaction_id"`
	MerchantID string  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Commision  float64 `json:"commission"`
	Fee        float64 `json:"fee"`
	CreatedAt  string  `json:"created_at"`
}
