package formatter

import "payment-gwf/entity"

type SimplePaymentResponse struct {
	TransactionID string     `json:"transaction_id"`
	OrderID       string     `json:"order_id"`
	VaNumbers     []VaNumber `json:"va_numbers"`
}

type VaNumber struct {
	Bank     string `json:"bank"`
	VaNumber string `json:"va_number"`
}

func FormatSimplePaymentResponse(midtrans *entity.DoPayment) SimplePaymentResponse {
	// Create an empty slice of VaNumber
	var vaNumbers []VaNumber

	// Iterate over VaNumbers in midtrans and add to vaNumbers slice
	for _, va := range midtrans.VaNumbers {
		vaNumbers = append(vaNumbers, VaNumber{
			Bank:     va.Bank,
			VaNumber: va.VaNumber,
		})
	}

	// Create and return the formatted response
	formatter := SimplePaymentResponse{
		TransactionID: midtrans.TransactionID,
		OrderID:       midtrans.OrderID,
		VaNumbers:     vaNumbers,
	}

	return formatter
}
