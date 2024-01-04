package adminHandler

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	// validate the struct
	validate := validator.New()
	if err := validate.Struct(Body); err != nil {
		res := response.ErrResponse{Response: "StructError", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Using a WaitGroup to wait for goroutine to finish
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Use a channel to signal the completion of the goroutines
	errCh := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		errCh <- adm.usecase.AdminSignup(ctx, Body)
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			res := response.ErrResponse{Response: "Error In Channel", Error: err.Error(), StatusCode: 500}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	case <-ctx.Done():
		// If the context times out, respond with an appropriate error
		res := response.ErrResponse{Response: nil, Error: "Request timed out", StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully created new admin", StatusCode: 200}
	c.JSON(http.StatusCreated, successRes)
}

func (adm *AdminHandler) AdminLogin(c *gin.Context) {
	var Body models.AdminLogin

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Body); err != nil {
		res := response.ErrResponse{Response: "Struct Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}


	ctx, cance := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cance()

	Token, err := adm.usecase.AdminLogin(ctx, Body)
	if err != nil {
		errRes := response.ErrResponse{Response: "Error In AdminLogin", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.LoginRes{TokenString: Token, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (adm *AdminHandler) DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Query Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = adm.usecase.DeleteAdmin(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error in Admin Delete", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully deleted admin", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}
