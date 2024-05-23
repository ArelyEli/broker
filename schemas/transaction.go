package schemas

type CreateTransactionRequest struct {
	Amount     float64 `json:"amount" binding:"required" required:"$field is required"`
	BusinessID uint    `json:"businessID" binding:"required" required:"$field is required"`
}
