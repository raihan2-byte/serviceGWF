package input

// type CreateOrder struct {
// 	// UserID  int   `json:"userID"`
// 	CartIDs []int `json:"cartIDs" binding:"required"`
// }

type CreateOrder struct {
	CartIDs []int `json:"cart_ids" binding:"required"`
	// Destination     string `json:"destination" binding:"required"`
	// Courier         string `json:"courier" binding:"required"`
	// ShippingPackage string `json:"shipping_package" binding:"required"`
	// HomeAddress     string `json:"home_address" binding:"required"`
}

type CreateOrderDetails struct {
	// CartIDs         []int  `json:"cart_ids" binding:"required"`
	Destination     string `json:"destination" binding:"required"`
	Courier         string `json:"courier" binding:"required"`
	ShippingPackage string `json:"shipping_package" binding:"required"`
	HomeAddress     string `json:"home_address" binding:"required"`
}
