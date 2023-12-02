package proiderHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type ProfileHandler struct {
	usecase interfaces.ProfileUsecase
}

func NewProfileHandler(use interfaces.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		usecase: use,
	}
}

func (p *ProfileHandler) AddService(c *gin.Context) {

	user_id:=c.GetInt("Uid")

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddService(user_id, service_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: "successfully added service", StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (p *ProfileHandler) DeleteService(c *gin.Context) {
	user_id:=c.GetInt("Uid")

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.DeleteService(user_id, service_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Data: "successfully removed service", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) AddPreferredWorkingLocation(c *gin.Context) {

	user_id:=c.GetInt("Uid")

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddPreferredWorkingLocation(user_id, district_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: "successfully added location to preferred locations", StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (p *ProfileHandler) RemovePreferredLocation(c *gin.Context) {
	user_id:=c.GetInt("Uid")

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.RemovePreferredLocation(user_id, district_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Data: "successfully removed location from preferred", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetSelectedServices(c *gin.Context) {
	user_id := c.GetInt("Uid")

	services, err := p.usecase.GetMyServices(user_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Data: services, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetAllPreferredLocations(c *gin.Context) {
	user_id:=c.GetInt("Uid")

	locations, err := p.usecase.GetAllPreferredLocations(user_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Data: locations, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetDetailsOfProviders(c *gin.Context) {
	pro_id:=c.GetInt("Uid")

	details, err := p.usecase.GetDetailsOfProviders(pro_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Data: details,StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

