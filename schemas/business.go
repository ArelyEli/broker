package schemas

type CreateBusinessRequest struct {
	Name       string  `json:"name" binding:"required" required:"$field is required"`
	Commission float64 `json:"commission" binding:"required,gte=1,lte=100" required:"$field is required con valor entre 1 y 100" gte:"$field must be greater than or equal to 1" lte:"$field must be less than or equal to 100"`
}

type UpdateBusinessRequest struct {
	Name       string  `json:"name" binding:"omitempty"`
	Commission float64 `json:"commission" binding:"omitempty,gte=1,lte=100" gte:"$field must be greater than or equal to 1" lte:"$field must be less than or equal to 100"`
}

type BusinessResponse struct {
	ID         uint    `json:"merchant_id"`
	Name       string  `json:"merchant_name"`
	Commission float64 `json:"commission"`
	CreatedAt  string  `json:"created_at"`
	UpdateAt   string  `json:"updated_at"`
}
