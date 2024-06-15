package formatter

import (
	"payment-gwf/entity"
	"time"
)

type StatusEkspedisi struct {
	ID        int       `json:"id"`
	ResiInfo  string    `json:"resi_info"`
	UserID    int       `json:"user_id"`
	OrderID   int       `json:"order_id"`
	OngkirID  int       `json:"ongkir_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PostStatusEkspedisiFormatter(status *entity.StatusEkspedisi) StatusEkspedisi {
	formatter := StatusEkspedisi{
		ID:        status.ID,
		ResiInfo:  status.ResiInfo,
		UserID:    status.UserID,
		OrderID:   status.OrderID,
		OngkirID:  status.OngkirID,
		CreatedAt: status.CreatedAt,
		UpdatedAt: status.UpdatedAt,
	}
	return formatter
}

type StatusEkspedisiGet struct {
	ID        int                   `json:"id"`
	ResiInfo  string                `json:"resi_info"`
	UserID    int                   `json:"user_id"`
	User      User                  `json:"user"`
	OrderID   int                   `json:"order_id"`
	OngkirID  int                   `json:"ongkir_id"`
	Order     GetOrder              `json:"order"`
	Ongkir    ApplyShippingResponse `json:"ongkir"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type GetOrder struct {
	ID          int `json:"id"`
	OngkirID    int `json:"ekspedisi_id"`
	TotalPrice  int `json:"total_price"`
	ShippingFee int `json:"shipping_fee"`
}

func GetStatusEkspedisiFormatter(status *entity.StatusEkspedisi) StatusEkspedisiGet {

	user := status.User

	userFormatter := User{
		Username: user.Username,
	}

	ongkir := status.Ongkir

	ongkirFormatter := ApplyShippingResponse{
		CityName:    ongkir.CityName,
		Province:    ongkir.Province,
		PostalCode:  ongkir.PostalCode,
		HomeAddress: ongkir.HomeAddress,
		Courier:     ongkir.Courier,
	}

	order := status.Order
	orderFormatter := GetOrder{
		ID:          status.ID,
		OngkirID:    order.OngkirID,
		TotalPrice:  order.TotalPrice,
		ShippingFee: order.ShippingFee,
	}

	formatter := StatusEkspedisiGet{
		ID:        status.ID,
		ResiInfo:  status.ResiInfo,
		UserID:    status.UserID,
		OrderID:   status.OrderID,
		OngkirID:  status.OngkirID,
		Order:     orderFormatter,
		Ongkir:    ongkirFormatter,
		User:      userFormatter,
		CreatedAt: status.CreatedAt,
		UpdatedAt: status.UpdatedAt,
	}
	return formatter
}

func FormatterGetAllStatusEkspedisi(status []*entity.StatusEkspedisi) []StatusEkspedisiGet {
	getAllStatus := []StatusEkspedisiGet{}

	for _, gets := range status {
		orderFormatter := GetStatusEkspedisiFormatter(gets)
		getAllStatus = append(getAllStatus, orderFormatter)
	}

	return getAllStatus
}

// package formatter

// import (
// 	"payment-gwf/entity"
// 	"time"
// )

// type StatusEkspedisi struct {
// 	ID        int       `json:"id"`
// 	ResiInfo  string    `json:"resi_info"`
// 	UserID    int       `json:"user_id"`
// 	OrderID   int       `json:"order_id"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// func PostStatusEkspedisiFormatter(status *entity.StatusEkspedisi) StatusEkspedisi {
// 	formatter := StatusEkspedisi{
// 		ID:        status.ID,
// 		ResiInfo:  status.ResiInfo,
// 		UserID:    status.UserID,
// 		OrderID:   status.OrderID,
// 		CreatedAt: status.CreatedAt,
// 		UpdatedAt: status.UpdatedAt,
// 	}
// 	return formatter
// }

// type StatusEkspedisiGet struct {
// 	ID        int       `json:"id"`
// 	ResiInfo  string    `json:"resi_info"`
// 	UserID    int       `json:"user_id"`
// 	User      User      `json:"user"`
// 	OrderID   int       `json:"order_id"`
// 	Order     GetOrder  `json:"order"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// type GetOrder struct {
// 	ID            int    `json:"id"`
// 	OngkirID      int    `json:"ekspedisi_id"`
// 	TotalPrice    int    `json:"total_price"`
// 	StatusPayment string `json:"status_payment"`
// 	ShippingFee   int    `json:"shipping_fee"`
// }

// func GetStatusEkspedisiFormatter(status *entity.StatusEkspedisi) StatusEkspedisiGet {

// 	user := status.User

// 	userFormatter := User{
// 		Username: user.Username,
// 	}

// 	order := status.Order

// 	orderFormatter := GetOrder{
// 		ID:            order.ID,
// 		OngkirID:      order.OngkirID,
// 		TotalPrice:    order.TotalPrice,
// 		StatusPayment: order.StatusPayment,
// 		ShippingFee:   order.ShippingFee,
// 	}

// 	formatter := StatusEkspedisiGet{
// 		ID:        status.ID,
// 		ResiInfo:  status.ResiInfo,
// 		UserID:    status.UserID,
// 		OrderID:   status.OrderID,
// 		Order:     orderFormatter,
// 		User:      userFormatter,
// 		CreatedAt: status.CreatedAt,
// 		UpdatedAt: status.UpdatedAt,
// 	}
// 	return formatter
// }

// func FormatterGetAllStatusEkspedisi(status []*entity.StatusEkspedisi) []StatusEkspedisiGet {
// 	getAllStatus := []StatusEkspedisiGet{}

// 	for _, gets := range status {
// 		orderFormatter := GetStatusEkspedisiFormatter(gets)
// 		getAllStatus = append(getAllStatus, orderFormatter)
// 	}

// 	return getAllStatus
// }
