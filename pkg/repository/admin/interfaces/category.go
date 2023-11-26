package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type CategoryRepo interface{
	CheckCategory(context.Context,string)(bool,error)
	CreateCategory(ctx context.Context,Cat string)(error)
	ListCatgory(context.Context ,models.PageNation)([]domain.Category,error)
	CheckCategoryById(context.Context,int)(bool,error)
	DeleteCategory(context.Context,int)(error)
}