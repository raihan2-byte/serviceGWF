package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var (
	listOfBank = map[string]midtrans.Bank{
		"bca":  midtrans.BankBca,
		"bri":  midtrans.BankBri,
		"cimb": midtrans.BankCimb,
		"bni":  midtrans.BankBni,
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
	HandleNotification(res *entity.MidtransNotificationRequest) error
	GetAllByUserID(ID int) ([]*entity.Payment, error)
	DoPayment(req input.SubmitPaymentRequest, orderID string, userID int) (*entity.DoPayment, error)
	FindStatus(orderID string) (*entity.Payment, error)
}

type servicePayment struct {
	repositoryPayment        repository.RepositoryPayment
	repositoryUser           repository.RepositoryUser
	repositoryOrder          repository.RepositoryOrder
	repositoryPaymentDetails repository.RepositoryPaymentDetails
	midtransGateway          *midtransGateway
}

func NewServicePayment(repositoryPayment repository.RepositoryPayment, repositoryUser repository.RepositoryUser, repositoryOrder repository.RepositoryOrder, repositoryPaymentDetails repository.RepositoryPaymentDetails, midtransGateway *midtransGateway) *servicePayment {
	return &servicePayment{
		repositoryPayment: repositoryPayment,
		repositoryUser:    repositoryUser,
		repositoryOrder:   repositoryOrder,
		// repositoryMidtrans: repositoryMidtrans,
		repositoryPaymentDetails: repositoryPaymentDetails,
		midtransGateway:          midtransGateway,
	}
}

func (s *servicePayment) HandleNotification(req *entity.MidtransNotificationRequest) error {
	findPayment, err := s.repositoryPayment.FindByTransactionID(req.TransactionID)
	if err != nil {
		return err
	}

	findPayment.StatusPayment = req.TransactionStatus

	_, err = s.repositoryPayment.Update(findPayment)
	if err != nil {
		return err
	}
	return nil
}

func (s *servicePayment) FindStatus(orderID string) (*entity.Payment, error) {
	get, err := s.repositoryPayment.FindByOrderId(orderID)
	if err != nil {
		return get, err
	}

	return get, nil
}

func (s *servicePayment) DoPayment(req input.SubmitPaymentRequest, orderID string, userID int) (*entity.DoPayment, error) {
	getOrder, err := s.repositoryOrder.FindById(orderID)
	if err != nil {
		return nil, fmt.Errorf("error fetching order: %w", err)
	}

	if getOrder.UserID != userID {
		return nil, fmt.Errorf("unauthorized user")
	}

	chosenBank, ok := listOfBank[req.BankTransfer]
	if !ok {
		return nil, fmt.Errorf("unsupported bank")
	}

	resp, err := s.midtransGateway.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  getOrder.ID,
			GrossAmt: int64(getOrder.TotalPrice),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: chosenBank,
		},
	})

	// if err != nil {
	// 	return nil, fmt.Errorf("error during charge transaction: %w", err)
	// }

	pay := &entity.Payment{
		StatusPayment: "pending",
		OrderID:       getOrder.ID,
		UserID:        getOrder.UserID,
		TransactionID: resp.TransactionID,
	}

	// mappedResp, err := mapChargeToResponse(resp)
	// if err != nil {
	// 	return nil, fmt.Errorf("error mapping charge response: %w", err)
	// }

	// _, err = s.repositoryPayment.SaveMidtrans(mappedResp)
	// if err != nil {
	// 	return nil, fmt.Errorf("error saving payment to Midtrans: %w", err)
	// }

	_, err = s.repositoryPayment.Save(pay)
	if err != nil {
		return nil, fmt.Errorf("error saving payment: %w", err)
	}

	return mapChargeToResponse(resp)
}

func mapChargeToResponse(resp *coreapi.ChargeResponse) (*entity.DoPayment, error) {
	if resp == nil {
		return nil, fmt.Errorf("nil response received")
	}

	var vaNumbers []entity.VaNumber
	for _, va := range resp.VaNumbers {
		vaNumbers = append(vaNumbers, entity.VaNumber{
			Bank:     va.Bank,
			VaNumber: va.VANumber,
		})
	}

	return &entity.DoPayment{
		TransactionID:     resp.TransactionID,
		OrderID:           resp.OrderID,
		GrossAmount:       resp.GrossAmount,
		PaymentType:       resp.PaymentType,
		TransactionStatus: resp.TransactionStatus,
		VaNumbers:         vaNumbers,
		TransactionTime:   resp.TransactionTime,
		MerchantID:        "G317569034",
	}, nil
}

func (s *servicePayment) GetAllPayment() ([]*entity.Payment, error) {
	address, err := s.repositoryPayment.FindAll()
	if err != nil {
		return address, err
	}
	return address, nil
}

// func (s *servicePayment) UpdatePayment(ID int) (*entity.Payment, error) {
// 	findPayment, err := s.repositoryPayment.FindById(ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	data := entity.Payment{}
// 	data.TransactionID =

// 	update, err := s.repositoryPayment.Update()
// }

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
