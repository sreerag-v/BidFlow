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
	page, err1 := strconv.Atoi(c.Query("page"))
	count, err2 := strconv.Atoi((c.Query("count")))

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

	leads, err := w.usecase.GetAllLeads(pro_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Err While Fetching Leads", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if leads == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Leads Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: leads, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) ViewLeads(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	page, err1 := strconv.Atoi(c.Query("page"))
	count, err2 := strconv.Atoi((c.Query("count")))

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

	lead, err := w.usecase.ViewLeads(work_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Viewing Leads", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if lead.User == "" {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Leads Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: lead, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) PlaceBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error  in Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	pro_id := c.GetInt("Uid")

	var body models.PlaceBid

	if err := c.Bind(&body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	body.WorkID = work_id
	body.ProID = pro_id

	err = w.usecase.PlaceBid(body)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Biding", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := response.SuccResponse{Response: "Bid has been placed successfully", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) ReplaceBidWithNewBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Pro_id := c.GetInt("Uid")

	var model models.PlaceBid

	if err := c.BindJSON(&model); err != nil {
		res := response.ErrResponse{Response: "Error While Bidnig Request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	model.WorkID = work_id
	model.ProID = Pro_id

	err = w.usecase.ReplaceBidWithNewBid(model)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Replacing Bid", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Old bid has been Replaced successfully with new bid", StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetAllOtherBidsOnTheLeads(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "input error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	bids, err := w.usecase.GetAllOtherBidsOnTheLeads(work_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error while fetching bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if len(bids) == 0 {
		res := response.ErrResponse{Error: "Bids Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)

}

func (w *ProWorkHandler) GetAllAcceptedBids(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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

	bids, err := w.usecase.GetAllAcceptedBids(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fething accepted bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if bids == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Accetped Bids Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetWorksOfAProvider(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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

	works, err := w.usecase.GetWorksOfAProvider(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fething the works", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if works == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: " Commited Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetAllOnGoingWorks(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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

	works, err := w.usecase.GetAllOnGoingWorks(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geing AllGoing Works", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if works == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "OnGoing Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (w *ProWorkHandler) GetCompletedWorks(c *gin.Context) {
	provider_id := c.GetInt("Uid")
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
	works, err := w.usecase.GetCompletedWorks(provider_id, pagenation)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Usecase", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if works == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Completed Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}
