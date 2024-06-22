package formatter

import (
	"payment-gwf/entity"
	"time"
)

type PaymentDonation struct {
	ID             int          `json:"id"`
	StatusPayment  string       `json:"status_payment"`
	MakeDonationID string       `json:"make_donation_id"`
	MakeDonation   MakeDonation `json:"make_donation"`
	UserID         int          `json:"user_id"`
	User           User         `json:"user"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type MakeDonation struct {
	ID      string `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	Message string `json:"message"`
}

func FormatterPaymentDonation(paymentDonation *entity.PaymentDonation) PaymentDonation {

	user := paymentDonation.User

	userFormatter := User{
		ID:       user.ID,
		Username: user.Username,
	}

	makeDonation := paymentDonation.MakeDonation

	makeDonationFormatter := MakeDonation{
		ID:      makeDonation.ID,
		UserID:  makeDonation.UserID,
		Name:    makeDonation.Name,
		Amount:  makeDonation.Amount,
		Message: makeDonation.Message,
	}

	formatter := PaymentDonation{
		ID:             paymentDonation.ID,
		StatusPayment:  paymentDonation.StatusPayment,
		MakeDonationID: paymentDonation.MakeDonationID,
		MakeDonation:   makeDonationFormatter,
		UserID:         paymentDonation.UserID,
		User:           userFormatter,
		CreatedAt:      paymentDonation.CreatedAt,
		UpdatedAt:      paymentDonation.UpdatedAt,
	}

	return formatter
}

func FormatterGetAllPaymentDonation(payment []*entity.PaymentDonation) []PaymentDonation {
	getAllPaymentDonation := []PaymentDonation{}

	for _, gets := range payment {
		paymentDonationFormatter := FormatterPaymentDonation(gets)
		getAllPaymentDonation = append(getAllPaymentDonation, paymentDonationFormatter)
	}

	return getAllPaymentDonation
}
