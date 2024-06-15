package handler

import (
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type addressHandler struct {
	addressService service.ServiceAddress
	authService    auth.Service
}

func NewAddressHandler(addressService service.ServiceAddress, authService auth.Service) *addressHandler {
	return &addressHandler{addressService, authService}
}

func (h *addressHandler) CreateAddress(c *gin.Context) {
	var input input.InputAddressBuyer

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	newProduct, err := h.addressService.CreateAddress(input, getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := (newProduct)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *addressHandler) GetAllAddress(c *gin.Context) {
	products, err := h.addressService.GetAllAddress()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get address")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *addressHandler) GetAddressByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.addressService.GetAddressByID(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get address")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *addressHandler) DeleteAddress(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.addressService.DeleteAddress(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get address")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}
