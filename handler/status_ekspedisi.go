package handler

import (
	"net/http"
	"payment-gwf/entity"
	"payment-gwf/formatter"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type statusEkspedisiHandler struct {
	statusEkspedisiService service.ServiceStatusEkspedisi
}

func NewStatusEkspedisiHandler(statusEkspedisiService service.ServiceStatusEkspedisi) *statusEkspedisiHandler {
	return &statusEkspedisiHandler{statusEkspedisiService}
}

func (h *statusEkspedisiHandler) CreateStatusEkspedisi(c *gin.Context) {
	var input input.InputStatusEkspedisi
	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.Param("user_id")
	paramUser, _ := strconv.Atoi(user)

	order := c.Param("order_id")
	paramOrder, _ := strconv.Atoi(order)

	create, err := h.statusEkspedisiService.CreateStatusEkspedisi(input, paramOrder, paramUser)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.PostStatusEkspedisiFormatter(create))
	c.JSON(http.StatusOK, response)
}

func (h *statusEkspedisiHandler) GetAllStatusEkspedisi(c *gin.Context) {
	get, err := h.statusEkspedisiService.GetAllStatusEkspedisi()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllStatusEkspedisi(get))
	c.JSON(http.StatusOK, response)
}

func (h *statusEkspedisiHandler) GetAllStatusEkspedisiByUserId(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	get, err := h.statusEkspedisiService.GetStatusEkspedisiByUser(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllStatusEkspedisi(get))
	c.JSON(http.StatusOK, response)
}

func (h *statusEkspedisiHandler) UpdateStatusEkspedisi(c *gin.Context) {
	param := c.Param("id")
	paramStatusEkspedisi, _ := strconv.Atoi(param)

	var input input.InputStatusEkspedisi
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	update, err := h.statusEkspedisiService.UpdateStatusEkspedisi(paramStatusEkspedisi, input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.PostStatusEkspedisiFormatter(update))
	c.JSON(http.StatusOK, response)
}

func (h *statusEkspedisiHandler) DeleteStatusEkspedisi(c *gin.Context) {
	param := c.Param("id")
	paramStatusEkspedisi, _ := strconv.Atoi(param)

	_, err := h.statusEkspedisiService.DeleteStatusEkspedisi(paramStatusEkspedisi)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "status ekspedisi has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
