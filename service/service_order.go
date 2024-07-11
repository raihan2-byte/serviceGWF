package service

import (
	"errors"
	"os"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
	"time"

	"github.com/google/uuid"
)

type ServiceOrder interface {
	// CreateOrder(userID int, inputOrder input.CreateOrder) (*entity.Order, error)
	CreateOrders(userID int, inputOrder input.CreateOrder) (*entity.Order, error)
	GetOrderHistoryByUserID(userID int) ([]*entity.Order, error)
	GetAllOrderHistory() ([]*entity.Order, error)
	CreateOrderDetails(inputOrder input.CreateOrderDetails, orderID string) (*entity.Order, error)
	FindOrderByID(userID int, orderID string) (*entity.Order, error)
	// GetAllCartByUserId(userID int) ([]*entity.Order, error)
	// UpdatedCart(getCartID int, getUserID int, Qty input.InputCart) (*entity.Order, error)
	// DeleteCart(userID int, cartID input.GetID) (*entity.Order, error)
}

type service_order struct {
	repository_order   repository.RepositoryOrder
	repository_cart    repository.RepositoryCart
	repository_product repository.RepositoryProduct
	repository_user    repository.RepositoryUser
	repository_ongkir  repository.RepositoryRajaOngkir
	repositoryPayment  repository.RepositoryPayment
	serviceRajaOngkir  *serviceRajaOngkir
}

func NewServiceOrder(repository_order repository.RepositoryOrder, repository_cart repository.RepositoryCart, repository_product repository.RepositoryProduct, repository_user repository.RepositoryUser, repository_ongkir repository.RepositoryRajaOngkir, repositoryPayment repository.RepositoryPayment, serviceRajaOngkir *serviceRajaOngkir) *service_order {
	return &service_order{repository_order, repository_cart, repository_product, repository_user, repository_ongkir, repositoryPayment, serviceRajaOngkir}
}

// func (s *service_order) CreateOrders(userID int, inputOrder input.CreateOrder) (*entity.Order, error) {
// 	// Get user data by ID
// 	getUser, err := s.repository_user.FindById(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get cart items based on list of IDs
// 	getAllCart, err := s.repository_cart.FindByIds(inputOrder.CartIDs)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check if cart is empty
// 	if len(getAllCart) == 0 {
// 		return nil, errors.New("cart empty")
// 	}

// 	// Check if all cart items belong to the same user
// 	for _, cartItem := range getAllCart {
// 		if cartItem.UserID != userID {
// 			return nil, errors.New("cart items do not belong to the user")
// 		}
// 	}

// 	// Initialize variable for total price
// 	var totalPrice int

// 	// Create slice to store order items
// 	var orderItems []entity.OrderItem

// 	// Loop through all cart items to calculate total price
// 	for _, item := range getAllCart {
// 		// Get product information based on product ID in cart
// 		product, err := s.repository_product.FindById(item.ProductID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Check product stock
// 		if product.Stock < item.Quantity {
// 			return nil, errors.New("insufficient stock for product: " + product.Name)
// 		}

// 		// Add product price to total price
// 		totalPrice += product.Price * item.Quantity

// 		// Create OrderItem entity
// 		orderItem := entity.OrderItem{
// 			ProductID: product.ID,
// 			Quantity:  item.Quantity,
// 			Price:     product.Price,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		// Add to orderItems slice
// 		orderItems = append(orderItems, orderItem)
// 	}

// 	// Calculate shipping fee using ApplyShipping function
// 	shippingParams := entity.ShippingFeeParams{
// 		Origin:      os.Getenv("API_ONGKIR_ORIGIN"),
// 		Destination: inputOrder.Destination, // Ensure inputOrder has this field
// 		Weight:      1000,                   // Adjust as needed
// 		Courier:     inputOrder.Courier,     // Ensure inputOrder has this field
// 		HomeAddress: inputOrder.HomeAddress,
// 	}
// 	shippingResponse, err := s.serviceRajaOngkir.ApplyShipping(shippingParams, inputOrder.ShippingPackage, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Add shipping fee to total price
// 	totalPrice += shippingResponse.ShippingFee

// 	// Create new order entity
// 	order := &entity.Order{
// 		ID:          uuid.New().String(),
// 		UserID:      getUser.ID,
// 		TotalPrice:  totalPrice,
// 		ShippingFee: shippingResponse.ShippingFee,
// 		OngkirID:    shippingResponse.ID,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	// Save order to database
// 	newOrder, err := s.repository_order.Save(order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Loop to save each order item and reduce product stock
// 	for _, item := range orderItems {
// 		product, err := s.repository_product.FindById(item.ProductID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Reduce product stock
// 		product.Stock -= item.Quantity
// 		_, err = s.repository_product.Update(product)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Set OrderID for OrderItem and save to database
// 		item.OrderID = newOrder.ID
// 		_, err = s.repository_order.SaveOrderItem(item)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// Clear cart items based on IDs after order is created
// 	err = s.repository_cart.ClearCartByIds(inputOrder.CartIDs)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newOrder, nil
// }

func (s *service_order) FindOrderByID(userID int, orderID string) (*entity.Order, error) {
	getUser, err := s.repository_user.FindById(userID)
	if err != nil {
		return nil, err
	}

	getOrder, err := s.repository_order.FindById(orderID)
	if err != nil {
		return nil, err
	}

	if getUser.ID != getOrder.UserID {
		return nil, errors.New("user account user have this order")
	}

	return getOrder, nil
}

func (s *service_order) CreateOrders(userID int, inputOrder input.CreateOrder) (*entity.Order, error) {
	// Get user data by ID
	getUser, err := s.repository_user.FindById(userID)
	if err != nil {
		return nil, err
	}

	// Get cart items based on list of IDs
	getAllCart, err := s.repository_cart.FindByIds(inputOrder.CartIDs)
	if err != nil {
		return nil, err
	}

	// Check if cart is empty
	if len(getAllCart) == 0 {
		return nil, errors.New("cart empty")
	}

	// Check if all cart items belong to the same user
	for _, cartItem := range getAllCart {
		if cartItem.UserID != userID {
			return nil, errors.New("cart items do not belong to the user")
		}
	}

	// Initialize variable for total price
	var totalPrice int

	// Create slice to store order items
	var orderItems []entity.OrderItem

	// Loop through all cart items to calculate total price
	for _, item := range getAllCart {
		// Get product information based on product ID in cart
		product, err := s.repository_product.FindById(item.ProductID)
		if err != nil {
			return nil, err
		}

		// Check product stock
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		// Add product price to total price
		totalPrice += product.Price * item.Quantity

		// Create OrderItem entity
		orderItem := entity.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Add to orderItems slice
		orderItems = append(orderItems, orderItem)
	}

	// Create new order entity
	order := &entity.Order{
		ID:          uuid.New().String(),
		UserID:      getUser.ID,
		TotalPrice:  totalPrice,
		ShippingFee: 0,
		OngkirID:    nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save order to database
	newOrder, err := s.repository_order.Save(order)
	if err != nil {
		return nil, err
	}

	// Loop to save each order item and reduce product stock
	for _, item := range orderItems {
		product, err := s.repository_product.FindById(item.ProductID)
		if err != nil {
			return nil, err
		}

		// Reduce product stock
		product.Stock -= item.Quantity
		_, err = s.repository_product.Update(product)
		if err != nil {
			return nil, err
		}

		// Set OrderID for OrderItem and save to database
		item.OrderID = newOrder.ID
		_, err = s.repository_order.SaveOrderItem(item)
		if err != nil {
			return nil, err
		}
	}

	// Clear cart items based on IDs after order is created
	err = s.repository_cart.ClearCartByIds(inputOrder.CartIDs)
	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *service_order) CreateOrderDetails(inputOrder input.CreateOrderDetails, orderID string) (*entity.Order, error) {
	// Find the existing order by ID
	order, err := s.repository_order.FindById(orderID)
	if err != nil {
		return nil, err
	}

	// Calculate shipping fee using ApplyShipping function
	shippingParams := entity.ShippingFeeParams{
		Origin:      os.Getenv("API_ONGKIR_ORIGIN"),
		Destination: inputOrder.Destination,
		Weight:      1000, // Adjust as needed
		Courier:     inputOrder.Courier,
		HomeAddress: inputOrder.HomeAddress,
	}
	shippingResponse, err := s.serviceRajaOngkir.ApplyShipping(shippingParams, inputOrder.ShippingPackage, order.UserID)
	if err != nil {
		return nil, err
	}

	// Add shipping fee to total price
	order.TotalPrice += shippingResponse.ShippingFee
	order.ShippingFee = shippingResponse.ShippingFee
	order.OngkirID = &shippingResponse.ID
	order.UpdatedAt = time.Now()

	// Update the order in the database
	updatedOrder, err := s.repository_order.Update(order)
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

func (s *service_order) GetOrderHistoryByUserID(userID int) ([]*entity.Order, error) {
	getUserID, err := s.repository_user.FindById(userID)
	if err != nil {
		return nil, err
	}
	get, err := s.repository_order.FindAllByUserID(getUserID.ID)
	if err != nil {
		return get, err
	}

	return get, nil
}

func (s *service_order) GetAllOrderHistory() ([]*entity.Order, error) {

	get, err := s.repository_order.FindAll()
	if err != nil {
		return get, err
	}

	return get, nil
}
