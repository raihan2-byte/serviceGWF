package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
)

type ServiceStatusEkspedisi interface {
	CreateStatusEkspedisi(input input.InputStatusEkspedisi, orderID int, userID int) (*entity.StatusEkspedisi, error)
	GetAllStatusEkspedisi() ([]*entity.StatusEkspedisi, error)
	GetStatusEkspedisiByUser(ID int) ([]*entity.StatusEkspedisi, error)
	DeleteStatusEkspedisi(ID int) (*entity.StatusEkspedisi, error)
	UpdateStatusEkspedisi(ID int, input input.InputStatusEkspedisi) (*entity.StatusEkspedisi, error)
	// GetProductByCategory(ID int) ([]*entity.StatusEkspedisi, error)
}

type serviceStatusEkspedisi struct {
	repositoryStatusEkspedisi repository.RepositoryStatusEkspedisi
	repositoryOrder           repository.RepositoryOrder
	repositoryUser            repository.RepositoryUser
}

func NewServiceStatusEkspedisi(repositoryStatusEkspedisi repository.RepositoryStatusEkspedisi, repositoryOrder repository.RepositoryOrder, repositoryUser repository.RepositoryUser) *serviceStatusEkspedisi {
	return &serviceStatusEkspedisi{repositoryStatusEkspedisi, repositoryOrder, repositoryUser}
}

func (s *serviceStatusEkspedisi) UpdateStatusEkspedisi(ID int, input input.InputStatusEkspedisi) (*entity.StatusEkspedisi, error) {
	findStatus, err := s.repositoryStatusEkspedisi.FindByID(ID)
	if err != nil {
		return findStatus, err
	}

	val := &entity.StatusEkspedisi{}

	findStatus.ResiInfo = input.ResiInfo
	val.OrderID = findStatus.OrderID
	val.UserID = findStatus.UserID
	val.OngkirID = findStatus.OngkirID

	update, err := s.repositoryStatusEkspedisi.Update(findStatus)
	if err != nil {
		return update, err
	}
	return update, nil
}

func (s *serviceStatusEkspedisi) CreateStatusEkspedisi(input input.InputStatusEkspedisi, orderID int, userID int) (*entity.StatusEkspedisi, error) {
	findOrder, err := s.repositoryOrder.FindById(orderID)
	if err != nil {
		return nil, err
	}

	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	if findOrder.Ongkir.ID == 0 {
		return nil, fmt.Errorf("invalid OngkirID: %d", findOrder.Ongkir.ID)
	}

	val := &entity.StatusEkspedisi{
		ResiInfo: input.ResiInfo,
		OrderID:  findOrder.ID,
		UserID:   findUser.ID,
		OngkirID: findOrder.Ongkir.ID,
	}

	create, err := s.repositoryStatusEkspedisi.Save(val)
	if err != nil {
		return &entity.StatusEkspedisi{}, err
	}
	return create, nil
}

func (s *serviceStatusEkspedisi) GetAllStatusEkspedisi() ([]*entity.StatusEkspedisi, error) {
	getAll, err := s.repositoryStatusEkspedisi.FindAll()
	if err != nil {
		return getAll, err
	}
	return getAll, nil
}

func (s *serviceStatusEkspedisi) GetStatusEkspedisiByUser(ID int) ([]*entity.StatusEkspedisi, error) {

	getUserID, err := s.repositoryUser.FindById(ID)
	if err != nil {
		return nil, err
	}
	get, err := s.repositoryStatusEkspedisi.FindAllByUserID(getUserID.ID)
	if err != nil {
		return get, err
	}

	return get, nil
}

func (s *serviceStatusEkspedisi) DeleteStatusEkspedisi(ID int) (*entity.StatusEkspedisi, error) {
	find, err := s.repositoryStatusEkspedisi.FindByID(ID)
	if err != nil {
		return find, err
	}

	delete, err := s.repositoryStatusEkspedisi.Delete(find)
	if err != nil {
		return delete, err
	}
	return delete, nil
}
