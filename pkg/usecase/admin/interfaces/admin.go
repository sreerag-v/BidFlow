package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type AdminUsecase interface{
	AdminSignup(context.Context,domain.Admin)(error)
	AdminLogin(context.Context,models.AdminLogin)(string,error)
	DeleteAdmin(context.Context ,int)error
}