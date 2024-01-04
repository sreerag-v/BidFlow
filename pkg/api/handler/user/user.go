package userHandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/smtp"
	"github.com/sreerag_v/BidFlow/pkg/twilio"
	"github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"github.com/sreerag_v/BidFlow/pkg/utils/response"
)

type UserHandler struct {
	Usecase interfaces.UserUsecase
}

func NewUserHandler(Usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		Usecase: Usecase,
	}
}

func (usr *UserHandler) SignUp(c *gin.Context) {
	var Body models.UserSignup

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error ", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Body); err != nil {
		res := response.ErrResponse{Response: "Struct Validation Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := usr.Usecase.SignUp(Body); err != nil {
		res := response.ErrResponse{Response: "Error From Creating User", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Signup completed successfully", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) Login(c *gin.Context) {
	var Body models.Login

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := usr.Usecase.Login(Body)
	if err != nil {
		res := response.ErrResponse{Response: "Error From UserLogin", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.LoginRes{TokenString: token, StatusCode: 201}
	c.JSON(http.StatusCreated, res)
}

// Twilio Otp Login
func (usr *UserHandler) OtpLogin(c *gin.Context) {
	var body models.OTPLoginStruct

	if err := c.ShouldBindJSON(&body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if body.Email == "" && body.Phone == "" && body.Name == "" {
		err := errors.New("enter atleast user_name or email or phone")
		res := response.ErrResponse{Response: "Struct Validation Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var user domain.User
	copier.Copy(&user, &body)

	us, err := usr.Usecase.OtpLogin(user)

	if err != nil {
		res := response.ErrResponse{Response: "Error From OtpLogin", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	log.Print("number:=", user.Phone)
	// send the otp
	_, err = twilio.TwillioOtpSent("+91" + user.Phone)
	if err != nil {
		res := response.ErrResponse{Response: "Faild to send otp", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: us.ID, StatusCode: 201}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) LoginOtpVerify(c *gin.Context) {
	var Body models.OTPVerifyStruct

	if err := c.ShouldBindJSON(&Body); err != nil {
		res := response.ErrResponse{Response:"Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var use = domain.User{
		ID:    int(Body.UserID),
		Email: Body.Email,
	}

	user, err := usr.Usecase.GetUserDetails(use)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Geting UserDetails", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// verify the otp
	err = twilio.TwilioVerifyOTP("+91"+user.Phone, Body.OTP)
	if err != nil {
		res := response.ErrResponse{Response: "Otp Verification Faild", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	token, err := usr.Usecase.GetJwtToken(user)
	if err != nil {
		res := response.ErrResponse{Response: "Faild to Genarate Token", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.LoginRes{TokenString: token, StatusCode: 201}
	c.JSON(http.StatusCreated, res)

}

func (usr *UserHandler) UserProfile(c *gin.Context) {
	userid := c.GetInt("id")

	user, err := usr.Usecase.UserProfile(c, userid)
	if err != nil {
		res := response.ErrResponse{Response: "Error From UserProfile", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := response.SuccResponse{Response: user, StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) UpdateProfile(c *gin.Context){
	var body models.UpdateUser

	if err:=c.Bind(&body);err!=nil{
		res := response.ErrResponse{Response: "Binding Error ", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Uid:=c.GetInt("id")

	err:=usr.Usecase.UpdateProfile(Uid,body)
	if err != nil {
		res := response.ErrResponse{Response: "Error From Updating Profile", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := response.SuccResponse{Response: "Profile Updated SucessFull", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) ForgottPassword(c *gin.Context){
	var body models.Forgott

	if err:=c.Bind(&body);err!=nil{
		res := response.ErrResponse{Response: "Binding Error ", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	otp:=smtp.VerifyOTP(body.Email)
	err:=usr.Usecase.ForgottPassword(body,otp)
	if err != nil {
		res := response.ErrResponse{Response: "error while forgetting password", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Otp Sent Successfull Check on Gmail", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) ChangePassword(c *gin.Context){
	var body models.ChangePassword

	if err:=c.Bind(&body);err!=nil{
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err:=usr.Usecase.ChangePassword(body)
	if err != nil {
		res := response.ErrResponse{Response: "error while changin password", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := response.SuccResponse{Response: "Password Change Successfull", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}