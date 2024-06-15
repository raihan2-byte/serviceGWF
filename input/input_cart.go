package input

type InputCart struct {
	Quantity int `json:"quantity" binding:"required"`
}

type GetID struct {
	ID int `uri:"id" binding:"required"`
}
