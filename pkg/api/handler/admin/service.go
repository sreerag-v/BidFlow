package adminHandler

import (
	"github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
)

type ServiceHandler struct{
	Usecase interfaces.ServiceUsecase
}

func NewServiceHandler(Usecase interfaces.ServiceUsecase)*ServiceHandler{
	return &ServiceHandler{
		Usecase: Usecase,
	}
}
