package handler

import (
	"net/http"
	"os"
	"payment-gwf/entity"
	"payment-gwf/helper"
	"payment-gwf/service"

	"github.com/gin-gonic/gin"
)

type ProvinceHandler struct {
	service        service.ServiceRajaOngkir
	serviceAddress service.ServiceAddress
}

func NewProvinceHandler(service service.ServiceRajaOngkir, serviceAddress service.ServiceAddress) *ProvinceHandler {
	return &ProvinceHandler{service: service, serviceAddress: serviceAddress}
}

// type ProvinceHandler struct {
// 	service        service.ServiceRajaOngkir
// 	serviceAddress service.ServiceAddress
// }

// func NewProvinceHandler(service service.ServiceRajaOngkir, serviceAddress service.ServiceAddress) *ProvinceHandler {
// 	return &ProvinceHandler{service: service, serviceAddress: serviceAddress}
// }

func (h *ProvinceHandler) GetProvinces(c *gin.Context) {
	provinces, err := h.service.GetProvince()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIresponse(http.StatusOK, provinces)

	c.JSON(http.StatusOK, response)
}

func (h *ProvinceHandler) GetCityByProvinceID(c *gin.Context) {

	param := c.Param("id")

	city, err := h.service.GetCityByProvinceID(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIresponse(http.StatusOK, city)

	c.JSON(http.StatusOK, response)
}

func (h *ProvinceHandler) CalculateShippingFee(c *gin.Context) {
	var params entity.ShippingFeeParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Set default values for origin and weight if not provided
	if params.Origin == "" {
		params.Origin = "39" // Default origin
	}
	if params.Weight == 0 {
		params.Weight = 100 // Default weight
	}

	if params.Destination == "" || params.Courier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	shippingFeeOptions, err := h.service.CalculateShippingFee(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIresponse(http.StatusOK, shippingFeeOptions)

	c.JSON(http.StatusOK, response)

}

func (h *ProvinceHandler) ApplyShipping1(c *gin.Context) {
	origin := os.Getenv("API_ONGKIR_ORIGIN")

	// Struct untuk mengikat data dari body request
	type ShippingRequest struct {
		CityID          string `json:"city_id" binding:"required"`
		Courier         string `json:"courier" binding:"required"`
		ShippingPackage string `json:"shipping_package" binding:"required"`
	}

	var req ShippingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	shippingFeeOptions, err := h.service.CalculateShippingFee(entity.ShippingFeeParams{
		Origin:      origin,
		Destination: req.CityID,
		Weight:      100,
		Courier:     req.Courier,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var selectedShipping entity.ShippingFeeOption
	found := false
	for _, shippingOption := range shippingFeeOptions {
		if shippingOption.Service == req.ShippingPackage {
			selectedShipping = shippingOption
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Shipping package not found"})
		return
	}

	type ApplyShippingResponse struct {
		ShippingFee int `json:"shipping_fee"`
		TotalWeight int `json:"total_weight"`
	}

	applyShippingResponse := ApplyShippingResponse{
		ShippingFee: selectedShipping.Fee,
		TotalWeight: 100,
	}

	response := helper.APIresponse(http.StatusOK, applyShippingResponse)
	c.JSON(http.StatusOK, response)
}

func (h *ProvinceHandler) ApplyShipping(c *gin.Context) {
	origin := os.Getenv("API_ONGKIR_ORIGIN")

	// Struct untuk mengikat data dari body request
	type ShippingRequest struct {
		CityID          string `json:"city_id" binding:"required"`
		Courier         string `json:"courier" binding:"required"`
		ShippingPackage string `json:"shipping_package" binding:"required"`
		HomeAddress     string `json:"home_address" binding:"required"`
	}

	var req ShippingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	// Panggil fungsi ApplyShipping dari service
	applyShippingResponse, err := h.service.ApplyShipping(entity.ShippingFeeParams{
		Origin:      origin,
		Destination: req.CityID,
		Weight:      100,
		Courier:     req.Courier,
		HomeAddress: req.HomeAddress,
	}, req.ShippingPackage, getUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIresponse(http.StatusOK, applyShippingResponse)
	c.JSON(http.StatusOK, response)
}
