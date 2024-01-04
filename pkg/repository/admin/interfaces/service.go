package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ServiceRepo interface {
	CheckIfServiceAlreadyExists(context.Context, string) (bool, error)
	AddServicesToACategory(context.Context, models.AddServicesToACategory) error
	GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error)
	GetAllServices(ctx context.Context,page models.PageNation) ([]domain.Profession, error)
	DeleteService(ctx context.Context, id int) error
	ReActivateService(ctx context.Context, id int) error 
}