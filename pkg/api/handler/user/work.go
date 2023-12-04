package userHandler

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type WorkHandler struct{
	usecase interfaces.WorkUsecase
}

func NewWorkHandler(use interfaces.WorkUsecase)*WorkHandler{
	return &WorkHandler{
		usecase: use,
	}
}

func (work *WorkHandler) ListNewOpening(c *gin.Context) {
	var model domain.Work
	if err := c.BindJSON(&model); err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("model", model)

	err := work.usecase.ListNewOpening(model)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.SuccResponse{Data: "Successfully listed opening",StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

// func (work *WorkHandler) GetAllListedWorks(c *gin.Context) {

// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	works, err := work.usecase.GetAllListedWorks(id)
// 	if err != nil {
// 		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 500}
// 		c.JSON(http.StatusInternalServerError, res)
// 		return
// 	}

// 	//return result
// 	res := response.SuccResponse{Data: works, StatusCode: 200}
// 	c.JSON(http.StatusCreated, res)

// }