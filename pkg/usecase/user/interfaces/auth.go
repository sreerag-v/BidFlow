package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)



type UserUsecase interface{
	SignUp(Body models.UserSignup)error
	Login(Body models.Login)(string,error)
	OtpLogin(body domain.User)(domain.User,error)
	GetJwtToken(domain.User)(string,error)
	GetUserDetails(domain.User)(domain.User,error)

	UserProfile(context.Context,int)([]models.UserDetails,error)
	UpdateProfile(int,models.UpdateUser)error

	ForgottPassword(models.Forgott,string)error
	ChangePassword(models.ChangePassword)error
}