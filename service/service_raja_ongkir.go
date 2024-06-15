package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"payment-gwf/entity"
	"payment-gwf/repository"
	"strconv"

	"github.com/joho/godotenv"
)

type ServiceRajaOngkir interface {
	GetProvince() ([]entity.Province, error)
	GetCityByProvinceID(provinceID string) ([]entity.City, error)
	CalculateShippingFee(shippingParam entity.ShippingFeeParams) ([]entity.ShippingFeeOption, error)
	ApplyShipping(params entity.ShippingFeeParams, shippingPackage string, userID int) (*entity.ApplyShippingResponse, error)
}

type serviceRajaOngkir struct {
	repositoryRajaOngkir repository.RepositoryRajaOngkir
	repositoryAddress    repository.RepositoryAddress
	repositoryUser       repository.RepositoryUser
}

func NewServiceRajaOngkir(repositoryRajaOngkir repository.RepositoryRajaOngkir, repositoryAddress repository.RepositoryAddress, repositoryUser repository.RepositoryUser) *serviceRajaOngkir {
	return &serviceRajaOngkir{repositoryRajaOngkir, repositoryAddress, repositoryUser}
}

func (s *serviceRajaOngkir) GetProvince() ([]entity.Province, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}

	baseURL := os.Getenv("API_ONGKIR_BASE_URL")
	apiKey := os.Getenv("API_ONGKIR_KEY")

	if baseURL == "" || apiKey == "" {
		return nil, errors.New("API_ONGKIR_BASE_URL or API_ONGKIR_KEY is not set")
	}

	response, err := http.Get(baseURL + "province?key=" + apiKey)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var provinceResponse entity.ProvinceResponse
	err = json.Unmarshal(body, &provinceResponse)
	if err != nil {
		return nil, err
	}

	return provinceResponse.ProvinceData.Result, nil
}

func (s *serviceRajaOngkir) GetCityByProvinceID(provinceID string) ([]entity.City, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}

	baseURL := os.Getenv("API_ONGKIR_BASE_URL")
	apiKey := os.Getenv("API_ONGKIR_KEY")

	if baseURL == "" || apiKey == "" {
		return nil, errors.New("API_ONGKIR_BASE_URL or API_ONGKIR_KEY is not set")
	}

	response, err := http.Get(baseURL + "city?key=" + apiKey + "&province=" + provinceID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cityResponse entity.CityResponse
	err = json.Unmarshal(body, &cityResponse)
	if err != nil {
		return nil, err
	}

	return cityResponse.CityData.Result, nil
}

func (s *serviceRajaOngkir) CalculateShippingFee(shippingParam entity.ShippingFeeParams) ([]entity.ShippingFeeOption, error) {
	if shippingParam.Origin == "" || shippingParam.Destination == "" || shippingParam.Weight <= 0 || shippingParam.Courier == "" {
		return nil, errors.New("unsupported params")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}

	apiBaseURL := os.Getenv("API_ONGKIR_BASE_URL")

	apiKey := os.Getenv("API_ONGKIR_KEY")
	if apiKey == "" {
		return nil, errors.New("API key is missing")
	}

	params := url.Values{}

	params.Add("key", apiKey)
	params.Add("origin", shippingParam.Origin)
	params.Add("destination", shippingParam.Destination)
	params.Add("weight", strconv.Itoa(shippingParam.Weight))
	params.Add("courier", shippingParam.Courier)

	response, err := http.PostForm(apiBaseURL+"cost", params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	ongkirResponse := entity.OngkirResponse{}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(body, &ongkirResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}

	var shippingFeeOptions []entity.ShippingFeeOption
	for _, result := range ongkirResponse.OngkirData.Result {
		for _, cost := range result.Costs {
			shippingFeeOptions = append(shippingFeeOptions, entity.ShippingFeeOption{
				Service:            cost.Service,
				Fee:                cost.Cost[0].Value,
				DestinationDetails: ongkirResponse.OngkirData.DestinationDetails,
			})
		}
	}
	return shippingFeeOptions, nil
}

func (s *serviceRajaOngkir) ApplyShipping(params entity.ShippingFeeParams, shippingPackage string, userID int) (*entity.ApplyShippingResponse, error) {

	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	shippingFeeOptions, err := s.CalculateShippingFee(params)
	if err != nil {
		return nil, err
	}

	var selectedShipping entity.ShippingFeeOption
	found := false
	for _, shippingOption := range shippingFeeOptions {
		if shippingOption.Service == shippingPackage {
			selectedShipping = shippingOption
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("shipping package not found")
	}

	applyShippingResponse := &entity.ApplyShippingResponse{
		ShippingFee: selectedShipping.Fee,
		TotalWeight: params.Weight,
		CityID:      selectedShipping.DestinationDetails.CityID,
		CityName:    selectedShipping.DestinationDetails.CityName,
		PostalCode:  selectedShipping.DestinationDetails.PostalCode,
		Province:    selectedShipping.DestinationDetails.Province,
		HomeAddress: params.HomeAddress,
		Courier:     params.Courier,
		UserID:      findUser.ID,
	}

	save, err := s.repositoryRajaOngkir.Save(applyShippingResponse)
	if err != nil {
		return nil, err
	}
	return save, nil

	// return applyShippingResponse, nil
}
