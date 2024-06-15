package service

import (
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
)

type ServiceAddress interface {
	CreateAddress(input input.InputAddressBuyer, userID int) (*entity.BuyerAddress, error)
	GetAllAddress() ([]*entity.BuyerAddress, error)
	GetAddressByID(ID int) (*entity.BuyerAddress, error)
	DeleteAddress(ID int) (*entity.BuyerAddress, error)
	// GetProductByCategory(ID int) ([]*entity.BuyerAddress, error)
}

type serviceAddress struct {
	repositoryAddress repository.RepositoryAddress
	repositoryUser    repository.RepositoryUser
}

func NewServiceAddress(repositoryAddress repository.RepositoryAddress, repositoryUser repository.RepositoryUser) *serviceAddress {
	return &serviceAddress{repositoryAddress, repositoryUser}
}

// func (s *serviceAddress) GetProductByCategory(ID int) ([]*entity.BuyerAddress, error) {
// 	address, err := s.repositoryAddress.FindAllAddressByCategory(ID)
// 	if err != nil {
// 		return address, err
// 	}
// 	return address, nil
// }

func (s *serviceAddress) GetAllAddress() ([]*entity.BuyerAddress, error) {

	address, err := s.repositoryAddress.FindAll()
	if err != nil {	
		return address, err
	}
	return address, nil
}

func (s *serviceAddress) CreateAddress(input input.InputAddressBuyer, userID int) (*entity.BuyerAddress, error) {
	address := &entity.BuyerAddress{}

	getUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	address.Province = input.Province
	address.City = input.City
	address.SubDistrict = input.SubDistrict
	address.HomeAddress = input.HomeAddress
	address.UserID = getUser.ID

	newAddress, err := s.repositoryAddress.Save(address)
	if err != nil {
		return newAddress, err
	}
	return newAddress, nil
}

func (s *serviceAddress) GetAddressByID(ID int) (*entity.BuyerAddress, error) {

	address, err := s.repositoryAddress.FindById(ID)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *serviceAddress) DeleteAddress(ID int) (*entity.BuyerAddress, error) {

	address, err := s.repositoryAddress.FindById(ID)
	if err != nil {
		return address, err
	}
	addressDel, err := s.repositoryAddress.Delete(address)

	if err != nil {
		return addressDel, err
	}
	return addressDel, nil

}
