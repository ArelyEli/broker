package schemas

type CreateBusinessRequest struct {
	Name       string  `json:"name" binding:"required" required:"$field is required"`
	Commission float64 `json:"commission" binding:"omitempty,gte=1,lte=100" gte:"$field must be greater than or equal to 1" lte:"$field must be less than or equal to 100"`
}
