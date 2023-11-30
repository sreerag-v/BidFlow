package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ServiceUsecase interface{
	AddServicesToACategory(context.Context, models.AddServicesToACategory) error
	GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error)
	DeleteService(ctx context.Context, id int) error 
	ReActivateService(ctx context.Context, id int) error
}