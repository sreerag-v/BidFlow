package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type RegionRepo interface{
	AddNewState(context.Context, string) error
	CheckStateExists(context.Context,string)(bool,error)
	ListStates(ctx context.Context, page models.PageNation)([]domain.State,error)
	CheckStateExistsByid(context.Context,int)(bool,error)
	DeleteState(context.Context, int) error

	CheckIfDistrictAlreadyExists(context.Context, string) (bool, error)
	AddNewDistrict(context.Context, models.AddNewDistrict) error
	GetDistrictsFromState(context.Context, int) ([]domain.District, error)
	DeleteDistrictFromState(context.Context, int) error
	CheckIfDistrictExistByid(context.Context,int)(bool,error)
}