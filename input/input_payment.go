package input

type SubmitPaymentRequest struct {
	// OrderID      int
	Amount       int64  `json:"amount"`
	BankTransfer string `json:"bank_transfer"`
	PaymentType  string `json:"payment_type"`
}
