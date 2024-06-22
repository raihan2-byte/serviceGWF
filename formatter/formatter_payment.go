package formatter

import (
	"payment-gwf/entity"
	"time"
)

type Payment struct {
	ID            int          `json:"id"`
	StatusPayment string       `json:"status_payment"`
	OrderID       string       `json:"order_id"`
	UserID        int          `json:"user_id"`
	User          User         `json:"user"`
	Order         OrderPayment `json:"order"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type OrderPayment struct {
	ID         string                       `json:"id"`
	TotalPrice int                          `json:"total_price"`
	Items      []Item                       `json:"items"`
	Ongkir     ApplyShippingResponsePayment `json:"ongkir"`
}

type Item struct {
	ProductID int            `json:"product_id"`
	Quantity  int            `json:"quantity"`
	Product   ProductPayment `json:"product"`
}

type ProductPayment struct {
	Name string `json:"name"`
}

type ApplyShippingResponsePayment struct {
	Courier     string `json:"courier"`
	HomeAddress string `json:"home_address"`
}

func FormatterPayment(payment *entity.Payment) Payment {
	orderItems := make([]Item, len(payment.Order.Items))
	for i, item := range payment.Order.Items {
		orderItems[i] = Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Product: ProductPayment{
				Name: item.Product.Name,
			},
		}
	}

	ongkir := payment.Order.Ongkir

	ongkirFormatter := ApplyShippingResponsePayment{
		Courier:     ongkir.Courier,
		HomeAddress: ongkir.HomeAddress,
	}

	return Payment{
		ID:            payment.ID,
		StatusPayment: payment.StatusPayment,
		OrderID:       payment.OrderID,
		UserID:        payment.UserID,
		User: User{
			ID:       payment.UserID,
			Username: payment.User.Username,
		},
		Order: OrderPayment{
			ID:         payment.Order.ID,
			TotalPrice: payment.Order.TotalPrice,
			Items:      orderItems,
			Ongkir:     ongkirFormatter,
		},
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}

func FormatterGetPayments(payments []*entity.Payment) []Payment {
	paymentsFormatter := []Payment{}

	for _, payment := range payments {
		paymentFormatter := FormatterPayment(payment)
		paymentsFormatter = append(paymentsFormatter, paymentFormatter)
	}

	return paymentsFormatter
}
