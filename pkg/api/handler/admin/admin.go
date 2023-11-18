package adminHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type AdminHandler struct {
	usecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		usecase: usecase,
	}
}

func (adm *AdminHandler) AdminSignup(c *gin.Context) {
	var Body domain.Admin

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := adm.usecase.AdminSignup(ctx, Body)

	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Data: "successfully created new admin", StatusCode: 200}
	c.JSON(http.StatusCreated, successRes)
}

func (adm *AdminHandler) AdminLogin(c *gin.Context) {
	var Body models.AdminLogin

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cance := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cance()

	Token, err := adm.usecase.AdminLogin(ctx, Body)
	if err != nil {
		errRes := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.SuccResponse{Data: Token, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (adm *AdminHandler) DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = adm.usecase.DeleteAdmin(ctx, id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Data: "successfully deleted admin", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}
