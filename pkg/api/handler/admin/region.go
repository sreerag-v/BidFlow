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

type RegionHandler struct {
	usecase interfaces.RegionUsecase
}

func NewRegionHandler(usecase interfaces.RegionUsecase) *RegionHandler {
	return &RegionHandler{
		usecase: usecase,
	}
}

func (reg *RegionHandler) AddNewState(c *gin.Context) {
	var Body models.AddNewState
	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err := reg.usecase.AddNewState(ctx, Body.State)
	if err!=nil{
		res := response.ErrResponse{Response: "Error In Adding State", Error: err.Error(), StatusCode: 500}
	c.JSON(http.StatusInternalServerError, res)
	return
	}

	successRes := response.SuccResponse{Response: "Successfully added new state", StatusCode: 201}
	c.JSON(http.StatusCreated, successRes)
}


func (reg *RegionHandler) GetStates(c *gin.Context){
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

	Cat, err := reg.usecase.ListStates(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error In List State", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if Cat == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "State Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}
	successRes := response.SuccResponse{Response: Cat, StatusCode: 201}
	c.JSON(http.StatusOK, successRes)
}

func (reg *RegionHandler) DeleteState(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("state_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err=reg.usecase.DeleteState(ctx,id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Delete State", Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Successfully made state inactive",StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}



func (reg *RegionHandler) AddNewDistrict(c *gin.Context){
	var region models.AddNewDistrict
	err := c.BindJSON(&region)
	if err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = reg.usecase.AddNewDistrict(ctx, region)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Add State", Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully added new District", StatusCode: 200}
	c.JSON(http.StatusCreated, successRes)
}

func (reg *RegionHandler) GetDistrictsFromState(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("state_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase get array
	districts, err := reg.usecase.GetDistrictsFromState(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In District From State", Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: districts, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (reg *RegionHandler) DeleteDistrictFromState(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err = reg.usecase.DeleteDistrictFromState(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error in Delete District", Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Successfully deleted district", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}
