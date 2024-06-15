package service

import (
	"errors"
	"os"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
)

type ServiceOrder interface {
	// CreateOrder(userID int, inputOrder input.CreateOrder) (*entity.Order, error)
	CreateOrders(userID int, inputOrder input.CreateOrder) (*entity.Order, error)
	GetOrderHistoryByUserID(userID int) ([]*entity.Order, error)
	GetAllOrderHistory() ([]*entity.Order, error)
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

func (s *service_order) CreateOrders(userID int, inputOrder input.CreateOrder) (*entity.Order, error) {
	// Dapatkan data pengguna berdasarkan ID
	getUser, err := s.repository_user.FindById(userID)
	if err != nil {
		return nil, err
	}

	// Dapatkan item dalam keranjang berdasarkan daftar ID
	getAllCart, err := s.repository_cart.FindByIds(inputOrder.CartIDs)
	if err != nil {
		return nil, err
	}

	// Periksa apakah keranjang kosong
	if len(getAllCart) == 0 {
		return nil, errors.New("cart empty")
	}

	// Periksa apakah semua item di keranjang milik pengguna yang sama
	for _, cartItem := range getAllCart {
		if cartItem.UserID != userID {
			return nil, errors.New("cart items do not belong to the user")
		}
	}

	// Inisialisasi variabel untuk total harga
	var totalPrice int

	// Buat slice untuk menyimpan item pesanan
	var orderItems []entity.OrderItem

	// Loop melalui semua item di keranjang untuk menghitung total harga
	for _, item := range getAllCart {
		// Dapatkan informasi produk berdasarkan ID produk di keranjang
		product, err := s.repository_product.FindById(item.ProductID)
		if err != nil {
			return nil, err
		}

		// Periksa stok produk
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		// Tambahkan harga produk ke total harga
		totalPrice += product.Price * item.Quantity

		// Buat entitas OrderItem
		orderItem := entity.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		// Tambahkan ke slice orderItems
		orderItems = append(orderItems, orderItem)
	}

	// Hitung biaya pengiriman menggunakan fungsi ApplyShipping
	shippingParams := entity.ShippingFeeParams{
		Origin:      os.Getenv("API_ONGKIR_ORIGIN"),
		Destination: inputOrder.Destination, // Pastikan inputOrder memiliki field ini
		Weight:      1000,                   // Sesuaikan dengan kebutuhan Anda
		Courier:     inputOrder.Courier,     // Pastikan inputOrder memiliki field ini
		HomeAddress: inputOrder.HomeAddress,
	}
	shippingResponse, err := s.serviceRajaOngkir.ApplyShipping(shippingParams, inputOrder.ShippingPackage, userID)
	if err != nil {
		return nil, err
	}

	// Tambahkan biaya pengiriman ke total harga
	totalPrice += shippingResponse.ShippingFee

	// Buat entitas pesanan baru
	order := &entity.Order{
		UserID:      getUser.ID,
		TotalPrice:  totalPrice,
		ShippingFee: shippingResponse.ShippingFee, // Tambahkan field ini
		OngkirID:    shippingResponse.ID,
		// StatusPayment: "pending", // Misalnya status awal adalah pending
		Items: orderItems,
	}

	// Simpan pesanan ke dalam database
	newOrder, err := s.repository_order.Save(order)
	if err != nil {
		return nil, err
	}

	// Loop untuk menyimpan setiap item pesanan dan mengurangi stok produk
	for _, item := range orderItems {
		product, err := s.repository_product.FindById(item.ProductID)
		if err != nil {
			return nil, err
		}

		// Kurangi stok produk
		product.Stock -= item.Quantity
		_, err = s.repository_product.Update(product)
		if err != nil {
			return nil, err
		}

		// Set OrderID untuk OrderItem dan simpan ke dalam database
		item.OrderID = newOrder.ID
		_, err = s.repository_order.SaveOrderItem(item)
		if err != nil {
			return nil, err
		}
	}

	// Hapus item dalam keranjang berdasarkan ID setelah pesanan dibuat
	err = s.repository_cart.ClearCartByIds(inputOrder.CartIDs)
	if err != nil {
		return nil, err
	}

	savePayment := &entity.Payment{}
	savePayment.StatusPayment = "pending"
	savePayment.OrderID = newOrder.ID
	savePayment.UserID = newOrder.UserID

	_, err = s.repositoryPayment.Save(savePayment)
	if err != nil {
		return nil, err
	}

	return newOrder, nil
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

// func (s *service_order) UpdatedCart(getCartID int, getUserID int, Qty input.InputCart) (*entity.Order, error) {
// 	cartID, err := s.repository_cart.FindById(getCartID)
// 	if err != nil {
// 		return cartID, err
// 	}

// 	userID, err := s.repository_user.FindById(getUserID)
// 	if err != nil {
// 		return &entity.Order{}, err
// 	}

// 	productID, err := s.repository_product.FindById(cartID.ProductID)
// 	if err != nil {
// 		return &entity.Order{}, err
// 	}

// 	totalPrice := Qty.Quantity * productID.Price

// 	cartID.Quantity = Qty.Quantity
// 	cartID.TotalPrice = totalPrice
// 	cartID.ProductID = productID.ID
// 	cartID.UserID = userID.ID

// 	updatedCart, err := s.repository_cart.Update(cartID)
// 	if err != nil {
// 		return updatedCart, err
// 	}

// 	return updatedCart, nil
// }

// func (s *service_order) DeleteCart(userID int, cartID input.GetID) (*entity.Order, error) {
// 	// Cek apakah cart dengan cartID tersebut milik user dengan userID yang sesuai
// 	getCart, err := s.repository_cart.FindById(cartID.ID)
// 	if err != nil {
// 		return getCart, err
// 	}

// 	getUser, err := s.repository_user.FindById(userID)
// 	if err != nil {
// 		return &entity.Order{}, err
// 	}

// 	if getCart.UserID != getUser.ID {
// 		return nil, errors.New("cart does not belong to current user")
// 	}

// 	// Hapus cart
// 	del, err := s.repository_cart.Delete(getCart)
// 	if err != nil {
// 		return del, err
// 	}

// 	return del, nil
// }
