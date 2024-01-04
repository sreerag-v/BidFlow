package userHandler

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type WorkHandler struct {
	usecase interfaces.WorkUsecase
}

func NewWorkHandler(use interfaces.WorkUsecase) *WorkHandler {
	return &WorkHandler{
		usecase: use,
	}
}

func (work *WorkHandler) ListNewOpening(c *gin.Context) {
	var model domain.ReqWork
	if err := c.BindJSON(&model); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("model", model)

	err := work.usecase.ListNewOpening(model)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Creatig New Opening", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: "Successfully listed opening", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (work *WorkHandler) AddImageOfWork(c *gin.Context) {
	imagepath, _ := c.FormFile("image")
	work_id, err := strconv.Atoi(c.PostForm("work_id"))
	if err != nil {
		res := response.ErrResponse{Response: "invalid input", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	extension := filepath.Ext(imagepath.Filename)
	image := uuid.New().String() + extension
	c.SaveUploadedFile(imagepath, "./public/images"+image)

	err = work.usecase.AddImageOfWork(image, work_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Adding image", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Successfully Added Image", StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}
func (work *WorkHandler) GetAllListedWorks(c *gin.Context) {

	id := c.GetInt("id")
	works, err := work.usecase.GetAllListedWorks(id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From listing Works", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if len(works) == 0 {
		res := response.ErrResponse{Response: "Works Not Listed", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllCompletedWorks(c *gin.Context) {

	id := c.GetInt("id")

	works, err := p.usecase.ListAllCompletedWorks(id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Listing Completed Works", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if len(works) == 0 {
		res := response.ErrResponse{Response: "No Works Are Completed", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllOngoingWorks(c *gin.Context) {

	id := c.GetInt("id")

	works, err := p.usecase.ListAllOngoingWorks(id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Listing OnGoing Works", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return

	}

	if len(works) == 0 {
		res := response.ErrResponse{Response: "No Works Are OnGoing", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: works, StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) WorkDetailsById(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error From Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Details, err := p.usecase.WorkDetailsById(work_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From WorkDetailsBy id", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if Details.WorkStatus == "" {
		res := response.ErrResponse{Response: "Works Not Listed", Error: "Works Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: Details, StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (p *WorkHandler) AcceptBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "error in request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	pro_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.ErrResponse{Response: "error in request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	bid_id, err := strconv.Atoi(c.Query("bid_id"))
	if err != nil {
		res := response.ErrResponse{Response: "error in request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	Uid := c.GetInt("id")

	err = p.usecase.AcceptBid(work_id, pro_id, bid_id, Uid)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Accepting Bid", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: "successfully accepted bid", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (p *WorkHandler) GetAllBids(c *gin.Context){
	count, err1 := strconv.Atoi(c.Query("count"))
	page, err2 := strconv.Atoi((c.Query("page")))

	err3 := errors.Join(err1, err2)

	if err3 != nil {
		res := response.ErrResponse{Response: "invalid input", Error: err3.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Uid:=c.GetInt("id")

	pagenation := models.PageNation{
		PageNumber: uint(page),
		Count:      uint(count),
	}
	bids,err:=p.usecase.GetAllBids(pagenation,Uid)

	if err != nil {
		res := response.ErrResponse{Response: "Err While Fetching Bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if bids == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: " Bids Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *WorkHandler) GetAllAcceptedBids(c *gin.Context){
	count, err1 := strconv.Atoi(c.Query("count"))
	page, err2 := strconv.Atoi((c.Query("page")))

	err3 := errors.Join(err1, err2)

	if err3 != nil {
		res := response.ErrResponse{Response: "invalid input", Error: err3.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Uid:=c.GetInt("id")

	pagenation := models.PageNation{
		PageNumber: uint(page),
		Count:      uint(count),
	}
	bids,err:=p.usecase.GetAllAcceptedBids(pagenation,Uid)

	if err != nil {
		res := response.ErrResponse{Response: "Err While Fetching Bids", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if bids == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: " Bids Not found ", StatusCode: 200}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.SuccResponse{Response: bids, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}

func (p *WorkHandler) AssignWorkToProvider(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "error in request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	pro_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.ErrResponse{Response: "error in request", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if work_id == 0 || pro_id == 0 {
		res := response.ErrResponse{Response: "checking error", Error: "invalid request parameters"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = p.usecase.AssignWorkToProvider(work_id, pro_id)
	if err != nil {
		res := response.ErrResponse{Response: "error in repo", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: "successfully assigned the work to provider", StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) MakeWorkAsCompleted(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error in Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.MakeWorkAsCompleted(work_id)
	if err != nil {
		res := response.ErrResponse{Response: "Error From MakeWorkCompleted", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: "Successfully completed work", StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) RateWork(c *gin.Context) {

	workID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var model models.RatingModel

	err = c.Bind(&model)
	if err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.RateWork(model, workID)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Repo", Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Response: "Rated successfully", StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (w *WorkHandler) RazorPaySent(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Param", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Uid := c.GetInt("id")
	razorid, amount, err := w.usecase.RazorPaySent(work_id, Uid)
	if err != nil {
		res := response.ErrResponse{Response: "Could Not Genarate OrderId", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.HTML(200, "app.html", gin.H{
		"userid":     Uid,
		"totalprice": amount,
		"paymentid":  razorid,
	})

}

func (w *WorkHandler) RazorPaySucess(c *gin.Context) {
	Uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.ErrResponse{Response: "Error In Query", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Oid := c.Query("order_id")
	Pid := c.Query("payment_id")
	Sig := c.Query("signature")
	total := c.Query("total")

	err = w.usecase.RazorPaySucess(Uid, Oid, Pid, Sig, total)
	if err != nil {
		res := response.ErrResponse{Response: "Error In Razor Pay Sucess", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(200, gin.H{
		"status":    true,
		"paymentid": Pid,
	})
}
