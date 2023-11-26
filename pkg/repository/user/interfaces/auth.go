package interfaces

import (
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
}