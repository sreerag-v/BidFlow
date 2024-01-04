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

type UserMgmtHandler struct {
	usecase interfaces.UserMgmtUsecase
}

func NewUserMgmtHandler(usecase interfaces.UserMgmtUsecase) *UserMgmtHandler {
	return &UserMgmtHandler{
		usecase: usecase,
	}
}

func (mg *UserMgmtHandler) GetProviders(c *gin.Context) {
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
	//call usecase get array
	providers, err := mg.usecase.GetProviders(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Get Providers", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if providers == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Providers Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}
	//give array
	successRes := response.SuccResponse{Response: providers, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) MakeProvidersVerified(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error in Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = mg.usecase.MakeProviderVerified(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Provider Verification", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Successfully Verified provider", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) RevokeVerification(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = mg.usecase.RevokeVerification(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Provider Varification Revoke", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Revoked verification of provider", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) GetUsers(c *gin.Context) {

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
	//call usecase get array
	users, err := mg.usecase.GetUsers(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Errorin Geting Users", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if users == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Providers Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}
	//[]
	successRes := response.SuccResponse{Response: users, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) BlockUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = mg.usecase.BlockUser(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error In BlockUser", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Blocked user", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) UnBlockUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = mg.usecase.UnBlockUser(ctx, id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From UnblockUser", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: "Unblocked user", StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}

func (mg *UserMgmtHandler) GetAllPendingVerifications(c *gin.Context) {
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
	//call usecase get array
	verifications, err := mg.usecase.GetAllPendingVerifications(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Provider Pending Verification", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if verifications == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found", Error: "Providers Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	//give array
	successRes := response.SuccResponse{Response: verifications, StatusCode: 200}
	c.JSON(http.StatusOK, successRes)
}
