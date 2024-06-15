package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var (
	listOfBank = map[string]midtrans.Bank{
		"bca":  midtrans.BankBca,
		"bri":  midtrans.BankBri,
		"cimb": midtrans.BankCimb,
	}
)

type midtransGateway struct {
	client *coreapi.Client
}

type Config struct {
	ServerKey string
}

func NewMidtransGateway(cfg *Config) (*midtransGateway, error) {
	client := coreapi.Client{}
	client.New(cfg.ServerKey, midtrans.Sandbox)
	return &midtransGateway{
		client: &client,
	}, nil
}

type ServicePayment interface {
	GetAllPayment() ([]*entity.Payment, error)
	GetPaymentByID(ID int) (*entity.Payment, error)
	DeletePayment(ID int) (*entity.Payment, error)
	GetAllByUserID(ID int) ([]*entity.Payment, error)
	DoPayment(req input.SubmitPaymentRequest, orderID int) (*entity.ResponDoPayment, error)
}

type servicePayment struct {
	repositoryPayment repository.RepositoryPayment
	repositoryUser    repository.RepositoryUser
	repositoryOrder   repository.RepositoryOrder
	midtransGateway   *midtransGateway
}

func NewServicePayment(repositoryPayment repository.RepositoryPayment, repositoryUser repository.RepositoryUser, repositoryOrder repository.RepositoryOrder, midtransGateway *midtransGateway) *servicePayment {
	return &servicePayment{
		repositoryPayment: repositoryPayment,
		repositoryUser:    repositoryUser,
		repositoryOrder:   repositoryOrder,
		midtransGateway:   midtransGateway,
	}
}

func (s *servicePayment) DoPayment(req input.SubmitPaymentRequest, orderID int) (*entity.ResponDoPayment, error) {
	getOrder, err := s.repositoryOrder.FindById((orderID))
	if err != nil {
		return nil, err
	}

	valGetOrderID := strconv.Itoa(getOrder.ID)

	chosenBank, ok := listOfBank[req.BankTransfer]
	if !ok {
		return nil, fmt.Errorf("unsupported bank")
	}

	resp, err := s.midtransGateway.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  valGetOrderID,
			GrossAmt: int64(getOrder.TotalPrice),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: chosenBank,
		},
	})

	if err != nil {
		return nil, err
	}

	return mapChargeToResponse(resp)
}

func mapChargeToResponse(resp *coreapi.ChargeResponse) (*entity.ResponDoPayment, error) {
	var vaNumbers []entity.VaNumber
	for _, va := range resp.VaNumbers {
		vaNumbers = append(vaNumbers, entity.VaNumber{
			Bank:     va.Bank,
			VaNumber: va.VANumber,
		})
	}

	return &entity.ResponDoPayment{
		TransactionID: resp.TransactionID,
		OrderID:       resp.OrderID,
		GrossAmount:   resp.GrossAmount,
		VaNumbers:     vaNumbers,
		MerchantID:    "G317569034",
	}, nil
}

func (s *servicePayment) GetAllPayment() ([]*entity.Payment, error) {
	address, err := s.repositoryPayment.FindAll()
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *servicePayment) GetPaymentByID(ID int) (*entity.Payment, error) {
	address, err := s.repositoryPayment.FindById(ID)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *servicePayment) GetAllByUserID(ID int) ([]*entity.Payment, error) {
	getUserID, err := s.repositoryUser.FindById(ID)
	if err != nil {
		return nil, err
	}
	get, err := s.repositoryPayment.FindAllByUserID(getUserID.ID)
	if err != nil {
		return get, err
	}
	return get, nil
}

func (s *servicePayment) DeletePayment(ID int) (*entity.Payment, error) {
	address, err := s.repositoryPayment.FindById(ID)
	if err != nil {
		return address, err
	}
	addressDel, err := s.repositoryPayment.Delete(address)
	if err != nil {
		return addressDel, err
	}
	return addressDel, nil
}
