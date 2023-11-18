package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)


type RegionUsecase interface{
	AddNewState(ctx context.Context,body string) error
	ListStates(context.Context,models.PageNation)([]domain.State,error)
	DeleteState(context.Context, int) error

	AddNewDistrict(context.Context, models.AddNewDistrict) error
	GetDistrictsFromState(context.Context, int) ([]domain.District, error)
	DeleteDistrictFromState(context.Context, int) error

}