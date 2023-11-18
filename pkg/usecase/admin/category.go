package adminUsecase

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)



type CategoryUsecase struct{
	Repo interfaces.CategoryRepo
}

func NewCategoryRepo(Repo interfaces.CategoryRepo)services.CategoryUsecase{
	return &CategoryUsecase{
		Repo: Repo,
	}
}

func (adm *CategoryUsecase) CreateCategory(ctx context.Context,Cat string)error{
	// chekc he category exist or not
	exists,err:=adm.Repo.CheckCategory(ctx,Cat)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("category already exists")
	}

	// create category
	err=adm.Repo.CreateCategory(ctx,Cat)
	if err != nil {
		return err
	}

	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (adm *CategoryUsecase) ListCatgory(ctx context.Context,page models.PageNation)([]domain.Category,error){
	Category,err:=adm.Repo.ListCatgory(ctx,page)
	if err != nil {
		return []domain.Category{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.Category{}, errors.New("request timeout")
	}
	return Category, nil
}

func (adm *CategoryUsecase)	DeleteCategory(ctx context.Context, id int)error{
	err:=adm.Repo.DeleteCategory(ctx,id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

