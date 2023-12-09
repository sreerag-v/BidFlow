package proiderHandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type ProWorkHandler struct {
	usecase interfaces.ProWorkUsecase
}

func NewProWorkHandler(use interfaces.ProWorkUsecase) *ProWorkHandler {
	return &ProWorkHandler{
		usecase: use,
	}
}

func (w *ProWorkHandler) GetAllLeads(c *gin.Context) {
	pro_id := c.GetInt("Uid")
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

	leads, err := w.usecase.GetAllLeads(pro_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: "Err While Fetching Leads", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if leads == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Leads Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Data: leads, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) ViewLeads(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
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

	lead, err := w.usecase.ViewLeads(work_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: lead, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) PlaceBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	pro_id := c.GetInt("Uid")

	var body models.PlaceBid

	if err := c.Bind(&body); err != nil {
		res := response.ErrResponse{Data: "binding error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	body.WorkID = work_id
	body.ProID = pro_id

	err = w.usecase.PlaceBid(body)
	if err != nil {
		res := response.ErrResponse{Data: "error while biding", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := response.SuccResponse{Data: "bid has been placed successfully", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) ReplaceBidWithNewBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Pro_id := c.GetInt("Uid")

	var model models.PlaceBid

	if err := c.BindJSON(&model); err != nil {
		res := response.ErrResponse{Data: "Error While Bidnig Request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	model.WorkID = work_id
	model.ProID = Pro_id

	err = w.usecase.ReplaceBidWithNewBid(model)
	if err != nil {
		res := response.ErrResponse{Data: "Error While Replacing Bid", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: "Old bid has been Replaced successfully with new bid", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetAllOtherBidsOnTheLeads(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Data: "input error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	bids, err := w.usecase.GetAllOtherBidsOnTheLeads(work_id)
	if err != nil {
		res := response.ErrResponse{Data: "Error while fetching bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (w *ProWorkHandler) GetAllAcceptedBids(c *gin.Context){
	provider_id := c.GetInt("Uid")
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

	bids,err:=w.usecase.GetAllAcceptedBids(provider_id,pagenation)
	if err != nil {
		res := response.ErrResponse{Data: "Error While Fething accepted bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if bids == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Data: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetWorksOfAProvider(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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

	works, err := w.usecase.GetWorksOfAProvider(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: "Error While Fething the works", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if works == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Data: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetAllOnGoingWorks(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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

	works, err := w.usecase.GetAllOnGoingWorks(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: "Error in Repo", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if works == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Data: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetCompletedWorks(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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
	works, err := w.usecase.GetCompletedWorks(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Data: "Error In Usecase", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if works == nil {
		res := response.ErrResponse{Data: "Go to Previous Page <.......", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Data: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}
