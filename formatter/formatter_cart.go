package formatter

import (
	"payment-gwf/entity"
	"time"
)

type Cart struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func AddToCartFormatter(cart *entity.Cart) Cart {
	formatter := Cart{
		ID:         cart.ID,
		UserID:     cart.UserID,
		ProductID:  cart.ProductID,
		Quantity:   cart.Quantity,
		TotalPrice: cart.TotalPrice,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}
	return formatter
}

type GetCart struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatterGetCartByUserID(cart *entity.Cart) GetCart {
	user := cart.User

	userFormatter := User{
		Username: user.Username,
	}

	formatter := GetCart{
		ID:         cart.ID,
		UserID:     cart.UserID,
		ProductID:  cart.ProductID,
		Quantity:   cart.Quantity,
		TotalPrice: cart.TotalPrice,
		User:       userFormatter,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}

	return formatter
}

func FormatterGetAllCartByUser(carts []*entity.Cart) []GetCart {
	getAllCart := []GetCart{}

	for _, gets := range carts {
		getAllFormatter := FormatterGetCartByUserID(gets)
		getAllCart = append(getAllCart, getAllFormatter)
	}

	return getAllCart
}
