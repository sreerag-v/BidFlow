package adminHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type ServiceHandler struct{
	Usecase interfaces.ServiceUsecase
}

func NewServiceHandler(Usecase interfaces.ServiceUsecase)*ServiceHandler{
	return &ServiceHandler{
		Usecase: Usecase,
	}
}

func (sr *ServiceHandler) AddServiceToCategory(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var service models.AddServicesToACategory
	err := c.BindJSON(&service)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = sr.Usecase.AddServicesToACategory(ctx, service)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.SuccResponse{Data: "successfully created service", StatusCode: 200}
	c.JSON(http.StatusCreated, successRes)
}

func (sr *ServiceHandler) GetServicesInACategory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase gest array
	services, err := sr.Usecase.GetServicesInACategory(ctx, id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Data: services, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (sr *ServiceHandler) DeleteService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = sr.Usecase.DeleteService(ctx, id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Data: "successfully deleted category", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (sr *ServiceHandler) ReActivateService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = sr.Usecase.ReActivateService(ctx, id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Data: "successfully Activated service", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}