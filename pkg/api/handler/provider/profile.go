package proiderHandler

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (p *ProfileHandler) AddProfileImage(c *gin.Context) {
	imagepath, _ := c.FormFile("image")
	Uid := c.GetInt("Uid")
	extension := filepath.Ext(imagepath.Filename)
	image := uuid.New().String() + extension
	c.SaveUploadedFile(imagepath, "./public/ProProfile"+image)
	err := p.usecase.AddProfileImage(image, Uid)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Adding Profileimage", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully Added ProfileImage", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (p *ProfileHandler) AddService(c *gin.Context) {

	user_id := c.GetInt("Uid")

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddService(user_id, service_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Adding Service TO Profile", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully added service", StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (p *ProfileHandler) DeleteService(c *gin.Context) {
	user_id := c.GetInt("Uid")

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.DeleteService(user_id, service_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Deleting Service from Profile", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully removed service", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) AddPreferredWorkingLocation(c *gin.Context) {

	user_id := c.GetInt("Uid")

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddPreferredWorkingLocation(user_id, district_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Adding Loaction", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully added location to preferred locations", StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (p *ProfileHandler) RemovePreferredLocation(c *gin.Context) {
	user_id := c.GetInt("Uid")

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.RemovePreferredLocation(user_id, district_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error Form Revmoving the location", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully removed location from preferred", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetSelectedServices(c *gin.Context) {
	user_id := c.GetInt("Uid")

	services, err := p.usecase.GetMyServices(user_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geting Service From Profile", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: services, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetAllPreferredLocations(c *gin.Context) {
	user_id := c.GetInt("Uid")

	locations, err := p.usecase.GetAllPreferredLocations(user_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geting Location", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: locations, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetDetailsOfProviders(c *gin.Context) {
	pro_id := c.GetInt("Uid")

	details, err := p.usecase.GetDetailsOfProviders(pro_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geting Provider Details", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: details, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *ProfileHandler) GetProDetails(c *gin.Context) {
	pro_id, err := strconv.Atoi(c.Query("Pid"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	details, err := p.usecase.GetProDetails(pro_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geting Provider Details", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.SuccResponse{Response: details, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}
