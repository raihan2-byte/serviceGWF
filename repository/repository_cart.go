package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryCart interface {
	FindAll() ([]*entity.Cart, error)
	Save(cart *entity.Cart) (*entity.Cart, error)
	FindById(ID int) (*entity.Cart, error)
	Update(cart *entity.Cart) (*entity.Cart, error)
	Delete(cart *entity.Cart) (*entity.Cart, error)
	ClearCartByUserID(userID int) error
	FindAllByUserID(ID int) ([]*entity.Cart, error)
	FindByIds(cartIDs []int) ([]*entity.Cart, error)
	ClearCartByIds(cartIDs []int) error
	FindByProductAndUser(productID int, userID int) (*entity.Cart, error)
}

type repository_cart struct {
	db *gorm.DB
}

func NewRepositoryCart(db *gorm.DB) *repository_cart {
	return &repository_cart{db}
}

func (r *repository_cart) FindAll() ([]*entity.Cart, error) {
	var carts []*entity.Cart
	err := r.db.Find(&carts).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (r *repository_cart) FindById(ID int) (*entity.Cart, error) {
	var carts *entity.Cart
	err := r.db.Where("id = ?", ID).Find(&carts).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (r *repository_cart) Save(cart *entity.Cart) (*entity.Cart, error) {

	err := r.db.Save(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository_cart) Update(cart *entity.Cart) (*entity.Cart, error) {

	err := r.db.Save(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository_cart) Delete(cart *entity.Cart) (*entity.Cart, error) {
	err := r.db.Delete(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *repository_cart) ClearCartByUserID(userID int) error {
	err := r.db.Where("user_id = ?", userID).Delete(&entity.Cart{}).Error
	return err
}

func (r *repository_cart) FindAllByUserID(ID int) ([]*entity.Cart, error) {
	var carts []*entity.Cart

	err := r.db.Preload("Product").Preload("User").Where("user_id = ?", ID).Find(&carts).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (r *repository_cart) FindByIds(cartIDs []int) ([]*entity.Cart, error) {
	var carts []*entity.Cart
	err := r.db.Where("id IN (?)", cartIDs).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *repository_cart) ClearCartByIds(cartIDs []int) error {
	err := r.db.Where("id IN (?)", cartIDs).Delete(&entity.Cart{}).Error
	return err
}

func (r *repository_cart) FindByProductAndUser(productID int, userID int) (*entity.Cart, error) {
	var carts *entity.Cart

	err := r.db.Where("product_id = ? ", productID).Where("user_id = ?", userID).Find(&carts).Error

	if err != nil {
		return carts, err
	}
	return carts, nil
}
