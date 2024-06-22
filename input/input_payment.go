package input

type SubmitPaymentRequest struct {
	// OrderID      int
	BankTransfer string `json:"bank_transfer"`
	PaymentType  string `json:"payment_type"`
}
