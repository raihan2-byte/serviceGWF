package input

type MakeDonationInput struct {
	Name    string `json:"name" binding:"required"`
	Amount  int    `json:"amount" binding:"required"`
	Message string `json:"message"`
}
