package service

import (
	"errors"
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
)

type ServiceCart interface {
	AddToCart(getIdProduct int, userID int, Qty input.InputCart) (*entity.Cart, error)
	GetAllCartByUserId(userID int) ([]*entity.Cart, error)
	UpdatedCart(getCartID int, getUserID int, Qty input.InputCart) (*entity.Cart, error)
	DeleteCart(userID int, cartID input.GetID) (*entity.Cart, error)
}

type service_cart struct {
	repository_cart    repository.RepositoryCart
	repository_product repository.RepositoryProduct
	repository_user    repository.RepositoryUser
}

func NewServiceCart(repository_cart repository.RepositoryCart, repository_product repository.RepositoryProduct, repository_user repository.RepositoryUser) *service_cart {
	return &service_cart{repository_cart, repository_product, repository_user}
}

func (s *service_cart) AddToCart(getIdProduct int, getUserID int, Qty input.InputCart) (*entity.Cart, error) {
	productID, err := s.repository_product.FindById(getIdProduct)
	if err != nil {
		return &entity.Cart{}, err
	}

	totalPrice := productID.Price * Qty.Quantity

	userID, err := s.repository_user.FindById(getUserID)
	if err != nil {
		return &entity.Cart{}, err
	}

	cart := entity.Cart{}

	cart.Quantity = Qty.Quantity
	cart.TotalPrice = totalPrice
	cart.ProductID = productID.ID
	cart.UserID = userID.ID

	exitingCart, err := s.repository_cart.FindByProductAndUser(getIdProduct, getUserID)
	if err != nil {
		return &entity.Cart{}, err
	}

	if exitingCart.ID != 0 {
		exitingCart.Quantity += Qty.Quantity
		exitingCart.TotalPrice = productID.Price * exitingCart.Quantity

		updatedCart, err := s.repository_cart.Update(exitingCart)
		if err != nil {
			return &entity.Cart{}, err
		}
		return updatedCart, nil
	}
	saveCart, err := s.repository_cart.Save(&cart)
	if err != nil {
		return saveCart, err
	}
	return saveCart, nil
}

func (s *service_cart) GetAllCartByUserId(userID int) ([]*entity.Cart, error) {
	getUserID, err := s.repository_user.FindById(userID)
	if err != nil {
		return nil, err
	}
	get, err := s.repository_cart.FindAllByUserID(getUserID.ID)

	if err != nil {
		return get, err
	}

	if len(get) == 0 {
		return nil, fmt.Errorf("no cart found for user with ID %d", userID)
	}

	return get, nil
}

func (s *service_cart) UpdatedCart(getCartID int, getUserID int, Qty input.InputCart) (*entity.Cart, error) {
	cartID, err := s.repository_cart.FindById(getCartID)
	if err != nil {
		return cartID, err
	}

	if cartID.ID == 0 {
		return nil, err
	}

	userID, err := s.repository_user.FindById(getUserID)
	if err != nil {
		return &entity.Cart{}, err
	}

	productID, err := s.repository_product.FindById(cartID.ProductID)
	if err != nil {
		return &entity.Cart{}, err
	}

	if cartID.UserID != userID.ID {
		return nil, errors.New("cart does not belong to current user")
	}

	totalPrice := Qty.Quantity * productID.Price

	cartID.Quantity = Qty.Quantity
	cartID.TotalPrice = totalPrice
	cartID.ProductID = productID.ID
	cartID.UserID = userID.ID

	updatedCart, err := s.repository_cart.Update(cartID)
	if err != nil {
		return updatedCart, err
	}

	return updatedCart, nil
}

func (s *service_cart) DeleteCart(userID int, cartID input.GetID) (*entity.Cart, error) {
	// Cek apakah cart dengan cartID tersebut milik user dengan userID yang sesuai
	getCart, err := s.repository_cart.FindById(cartID.ID)
	if err != nil {
		return getCart, err
	}

	getUser, err := s.repository_user.FindById(userID)
	if err != nil {
		return &entity.Cart{}, err
	}

	if getCart.UserID != getUser.ID {
		return nil, errors.New("cart does not belong to current user")
	}

	// Hapus cart
	del, err := s.repository_cart.Delete(getCart)
	if err != nil {
		return del, err
	}

	return del, nil
}
