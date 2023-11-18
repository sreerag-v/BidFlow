package adminUsecase

import (

	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
)

type ServiceUsecase struct{
	Repo interfaces.ServiceRepo
}

func NewServiceUsecase(Repo interfaces.ServiceRepo)services.ServiceUsecase{
	return &ServiceUsecase{
		Repo: Repo,
	}
}

