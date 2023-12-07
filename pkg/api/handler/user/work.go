package userHandler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
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
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("model", model)

	err := work.usecase.ListNewOpening(model)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: "Successfully listed opening", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (work *WorkHandler) AddImageOfWork(c *gin.Context){
	imagepath,_:=c.FormFile("image")
	work_id, err := strconv.Atoi(c.PostForm("work_id"))
	if err != nil {
		res := response.ErrResponse{Data: "invalid input", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	extension := filepath.Ext(imagepath.Filename)
	image := uuid.New().String() + extension
	c.SaveUploadedFile(imagepath, "./public/images"+image)

	err=work.usecase.AddImageOfWork(image,work_id)
	if err != nil {
		res := response.ErrResponse{Data: "Error While Adding image", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: "Successfully Added Image", StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}
func (work *WorkHandler) GetAllListedWorks(c *gin.Context) {

	id := c.GetInt("id")
	works, err := work.usecase.GetAllListedWorks(id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: works, StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllCompletedWorks(c *gin.Context) {

	id:=c.GetInt("id")

	works, err := p.usecase.ListAllCompletedWorks(id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: works,StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllOngoingWorks(c *gin.Context) {

	id:=c.GetInt("id")

	works, err := p.usecase.ListAllOngoingWorks(id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: works, StatusCode: 200}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) WorkDetailsById(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Details, err := p.usecase.WorkDetailsById(work_id)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: Details, StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}