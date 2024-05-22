package schemas

type CreateBusinessRequest struct {
	Name       string  `json:"name" binding:"required"`
	Commission float64 `json:"commission" binding:"omitempty,min=1,max=100"`
}
