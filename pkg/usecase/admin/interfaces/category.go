package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type CategoryUsecase interface{
	CreateCategory(ctx context.Context,Cat string)error
	ListCatgory(context.Context,models.PageNation)([]domain.Category,error)
	DeleteCategory(context.Context,int)error
}