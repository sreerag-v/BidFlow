package proiderHandler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type ProviderHandler struct{
	Usecase interfaces.ProviderUsecase
}

func NewProviderHandler(usecase interfaces.ProviderUsecase)*ProviderHandler{
	return &ProviderHandler{
		Usecase: usecase,
	}
}


func (pro *ProviderHandler) Register(c *gin.Context){
	name:=c.Request.FormValue("name")
	email:=c.Request.FormValue("email")
	password:=c.Request.FormValue("password")
	repassword:=c.Request.FormValue("repassword")
	phone:=c.Request.FormValue("phone")
	image,err:=c.FormFile("document")

	if err!=nil{
		res:=response.ErrResponse{Data: nil,Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest,res)
		return
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		res := response.ErrResponse{Data: nil, Error: "Invalid email format", StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var model models.ProviderRegister

	model.Name = name
	model.Email = email
	model.Password = password
	model.RePassword = repassword
	model.Phone = phone
	model.Document = image

	fmt.Println("model",model)

	// pass the values to usecase 
	err=pro.Usecase.Register(model)
	if err!=nil{
		res:=response.ErrResponse{Data: nil,Error: err.Error(),StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	//return result
	res := response.SuccResponse{Data: "Your request will be under inspection of admins", StatusCode: 201}
	c.JSON(http.StatusCreated, res)
}

func (pro *ProviderHandler) Login(c *gin.Context){
	var Body models.Login
	if err := c.BindJSON(&Body); err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token,err:=pro.Usecase.Login(Body)
	if err != nil {
		res := response.ErrResponse{Data: nil, Error: err.Error(),StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Data: token, StatusCode: 200}
	c.JSON(http.StatusOK, res)
}