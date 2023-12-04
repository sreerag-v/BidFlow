package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)


type UserRepo interface{
	CheckPhoneNumberExist(string)(bool,error)
	SignUp(models.UserSignup)error
	CheckUsername(string)(bool,error)
	GetUserDetails(string)(domain.User,error)
	GetUserDetailsById(uint)(domain.User,error)
	CheckUserBlockedOrNot(string)(domain.User,error)

	UserProfile(context.Context,int)([]models.UserDetails,error)
	UpdateProfile(int,models.UpdateUser)error

	FindUserByEmail(string)(domain.User,error)
	ForgottPassword(body models.Forgott,otp string)error
	ChangePassword(models.ChangePassword)error
}