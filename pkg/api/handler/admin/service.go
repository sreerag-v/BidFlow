package adminHandler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type ServiceHandler struct {
	Usecase interfaces.ServiceUsecase
}

func NewServiceHandler(Usecase interfaces.ServiceUsecase) *ServiceHandler {
	return &ServiceHandler{
		Usecase: Usecase,
	}
}

func (sr *ServiceHandler) AddServiceToCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var service models.AddServicesToACategory
	err := c.Bind(&service)
	if err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = sr.Usecase.AddServicesToACategory(ctx, service)
	if err != nil {
		res := response.ErrResponse{Response: "Error in Service Adding ", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully created service", StatusCode: 200}
	c.JSON(http.StatusCreated, successRes)
}

func (sr *ServiceHandler) GetServicesInACategory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase gest array
	services, err := sr.Usecase.GetServicesInACategory(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Listing Service ", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: services, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (sr *ServiceHandler) GetAllServices(c *gin.Context){
	count, err1 := strconv.Atoi(c.Query("count"))
	page, err2 := strconv.Atoi((c.Query("page")))

	err3 := errors.Join(err1, err2)

	if err3 != nil {
		res := response.ErrResponse{Response: "invalid input", Error: err3.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	pagenation := models.PageNation{
		PageNumber: uint(page),
		Count:      uint(count),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	Cat, err := sr.Usecase.GetAllServices(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error In List Category", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if Cat == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Services Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}
	successRes := response.SuccResponse{Response: Cat, StatusCode: 201}
	c.JSON(http.StatusOK, successRes)
}

func (sr *ServiceHandler) DeleteService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = sr.Usecase.DeleteService(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Delete Service", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Successfully deleted category", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (sr *ServiceHandler) ReActivateService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = sr.Usecase.ReActivateService(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In ReActive Service", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Successfully Activated service", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}
