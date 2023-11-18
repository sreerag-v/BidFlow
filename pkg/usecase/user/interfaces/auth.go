package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)



type UserUsecase interface{
	SignUp(Body models.UserSignup)error
	Login(Body models.Login)(string,error)
	OtpLogin(body domain.User)(domain.User,error)
	GetJwtToken(domain.User)(string,error)
	GetUserDetails(domain.User)(domain.User,error)
}