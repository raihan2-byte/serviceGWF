package formatter

import (
	"payment-gwf/entity"
	"time"
)

type Order struct {
	ID          string                `json:"id"`
	UserID      int                   `json:"user_id"`
	OngkirID    int                   `json:"ekspedisi_id"`
	TotalPrice  int                   `json:"total_price"`
	ShippingFee int                   `json:"shipping_fee"`
	Ongkir      ApplyShippingResponse `json:"ekspedisi"`
	Items       []OrderItems          `json:"items"`
	User        User                  `json:"user"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type OrderItems struct {
	Quantity int
	Products Products `json:"product"`
}

type Products struct {
	Name  string
	Price int
}

type User struct {
	ID       int
	Username string
}

type ApplyShippingResponse struct {
	CityName    string `json:"city_name"`
	PostalCode  string `json:"postal_code"`
	Province    string `json:"province"`
	HomeAddress string `json:"home_address" binding:"required"`
	Courier     string `json:"courier"`
}

type OrderItemPost struct {
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
	ProductID int `json:"product_id"`
}

func FormatterPostItem(item entity.OrderItem) OrderItemPost {
	formatterItem := OrderItemPost{
		Quantity:  item.Quantity,
		Price:     item.Price,
		ProductID: item.ProductID,
	}

	return formatterItem

}

type OrderPost struct {
	ID          string          `json:"id"`
	UserID      int             `json:"user_id"`
	OngkirID    int             `json:"ekspedisi_id"`
	TotalPrice  int             `json:"total_price"`
	ShippingFee int             `json:"shipping_fee"`
	Items       []OrderItemPost `json:"items"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func FormatterPostOrder(order *entity.Order) OrderPost {
	var formattedItems []OrderItemPost
	for _, item := range order.Items {
		formattedItem := FormatterPostItem(item)
		formattedItems = append(formattedItems, formattedItem)
	}

	formatter := OrderPost{
		ID:          order.ID,
		UserID:      order.UserID,
		OngkirID:    order.OngkirID,
		TotalPrice:  order.TotalPrice,
		ShippingFee: order.ShippingFee,
		Items:       formattedItems,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	return formatter
}

func FormatterItem(item entity.OrderItem) OrderItems {
	formatterItem := OrderItems{
		Quantity: item.Quantity,
	}

	product := item.Product

	productFormatter := Products{
		Name:  product.Name,
		Price: product.Price,
	}

	formatterItem.Products = productFormatter

	return formatterItem

}

func FormatterOrder(order *entity.Order) Order {
	var formattedItems []OrderItems
	for _, item := range order.Items {
		formattedItem := FormatterItem(item)
		formattedItems = append(formattedItems, formattedItem)
	}

	user := order.User

	userFormatter := User{
		ID:       user.ID,
		Username: user.Username,
	}

	ongkir := order.Ongkir

	ongkirFormatter := ApplyShippingResponse{
		CityName:    ongkir.CityName,
		Province:    ongkir.Province,
		PostalCode:  ongkir.PostalCode,
		HomeAddress: ongkir.HomeAddress,
		Courier:     ongkir.Courier,
	}

	formatter := Order{
		ID:          order.ID,
		UserID:      order.UserID,
		OngkirID:    order.OngkirID,
		TotalPrice:  order.TotalPrice,
		ShippingFee: order.ShippingFee,
		Items:       formattedItems,
		Ongkir:      ongkirFormatter,
		User:        userFormatter,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	return formatter
}

func FormatterGetAllOrder(order []*entity.Order) []Order {
	getAllOrder := []Order{}

	for _, gets := range order {
		orderFormatter := FormatterOrder(gets)
		getAllOrder = append(getAllOrder, orderFormatter)
	}

	return getAllOrder
}
