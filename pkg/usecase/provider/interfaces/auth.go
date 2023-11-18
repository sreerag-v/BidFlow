package interfaces

import "github.com/sreerag_v/BidFlow/pkg/utils/models"

type ProviderUsecase interface{
	Register(models.ProviderRegister)error
	Login(models.Login)(string,error)
}