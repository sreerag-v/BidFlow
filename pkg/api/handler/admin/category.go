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

type CategoryHandler struct {
	usecase interfaces.CategoryUsecase
}

func NewCategoryHandler(Usecase interfaces.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: Usecase,
	}
}

func (adm *CategoryHandler) CreateCategory(c *gin.Context) {
	var Category models.CreateCategory
	if err := c.Bind(&Category); err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err := adm.usecase.CreateCategory(ctx, Category.Category)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.SuccResponse{Data: "successfully created category", StatusCode: 201}
	c.JSON(http.StatusCreated, successRes)
}

func (adm *CategoryHandler) ListCatgory(c *gin.Context) {

	count, err1 := strconv.Atoi(c.Query("count"))
	page, err2 := strconv.Atoi((c.Query("page")))

	err3 := errors.Join(err1, err2)

	if err3 != nil {
		res := response.ErrResponse{Data: "invalid input", Error: err3.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	pagenation := models.PageNation{
		PageNumber: uint(page),
		Count:      uint(count),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	Cat, err := adm.usecase.ListCatgory(ctx, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if Cat == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Category Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}
	successRes := response.SuccResponse{Data: Cat, StatusCode: 201}
	c.JSON(http.StatusOK, successRes)
}

func (adm *CategoryHandler) DeleteCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = adm.usecase.DeleteCategory(ctx, id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Data: "successfully deleted category", StatusCode: 201}
	c.JSON(http.StatusOK, successRes)
}
