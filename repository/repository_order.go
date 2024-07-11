package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryOrder interface {
	FindAll() ([]*entity.Order, error)
	Save(order *entity.Order) (*entity.Order, error)
	SaveOrderItem(order entity.OrderItem) (entity.OrderItem, error)
	FindById(ID string) (*entity.Order, error)
	Update(order *entity.Order) (*entity.Order, error)
	Delete(order *entity.Order) (*entity.Order, error)
	FindAllByUserID(ID int) ([]*entity.Order, error)
	FindByProductAndUser(productID int, userID int) (*entity.Order, error)
}

type repository_order struct {
	db *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *repository_order {
	return &repository_order{db}
}

func (r *repository_order) FindAll() ([]*entity.Order, error) {
	var order []*entity.Order
	err := r.db.Preload("Items").Preload("Items.Product").Preload("Ongkir").Preload("User").Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) FindById(ID string) (*entity.Order, error) {
	var order *entity.Order
	err := r.db.Preload("Ongkir").Preload("Items").Preload("Items.Product").Where("id = ?", ID).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) Save(order *entity.Order) (*entity.Order, error) {

	err := r.db.Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) SaveOrderItem(order entity.OrderItem) (entity.OrderItem, error) {

	err := r.db.Preload("Items").Preload("Items.Product").Preload("User").Preload("Ongkir").Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) Update(order *entity.Order) (*entity.Order, error) {

	err := r.db.Preload("Items").Preload("Items.Product").Preload("User").Preload("Ongkir").Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) Delete(order *entity.Order) (*entity.Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository_order) FindAllByUserID(ID int) ([]*entity.Order, error) {
	var order []*entity.Order

	err := r.db.Preload("Items").Preload("Items.Product").Preload("User").Preload("Ongkir").Where("user_id = ?", ID).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository_order) FindByProductAndUser(productID int, userID int) (*entity.Order, error) {
	var order *entity.Order

	err := r.db.Where("product_id = ? ", productID).Where("user_id = ?", userID).Find(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}
