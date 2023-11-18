package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)



type ProviderRepo interface{
	CheckPhoneNumber(string)(bool,error)
	Register(model models.ProviderRegister)(int,error)
	UploadDoc(int,string)error

	CheckProExistOrNot(string)(bool,error)
	GetProDetails(string)(domain.Provider,error)
}
