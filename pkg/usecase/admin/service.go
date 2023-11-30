package adminUsecase

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ServiceUsecase struct{
	Repo interfaces.ServiceRepo
	UserMgmtRepo interfaces.UserMgmtRepo
	RegionRepo interfaces.RegionRepo
}

func NewServiceUsecase(Repo interfaces.ServiceRepo,	UserMgmtRepo interfaces.UserMgmtRepo,	RegionRepo interfaces.RegionRepo)services.ServiceUsecase{
	return &ServiceUsecase{
		Repo: Repo,
		UserMgmtRepo: UserMgmtRepo,
		RegionRepo: RegionRepo,
	}
}

func (sr *ServiceUsecase) AddServicesToACategory(ctx context.Context, service models.AddServicesToACategory) error {
	if err := ctx.Err(); err != nil {
		return errors.New("request timeout")
	}
	//check if already a category exists in same name
	exist, err := sr.Repo.CheckIfServiceAlreadyExists(ctx, service.ServiceName)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("service already exists")
	}
	//create new category
	err = sr.Repo.AddServicesToACategory(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (sr *ServiceUsecase) GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error) {

	service, err := sr.Repo.GetServicesInACategory(ctx, id)
	if err != nil {
		return []domain.Profession{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.Profession{}, errors.New("request timeout")
	}

	return service, nil
}

func (sr *ServiceUsecase) DeleteService(ctx context.Context, id int) error {

	err := sr.Repo.DeleteService(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (sr *ServiceUsecase) ReActivateService(ctx context.Context, id int) error {

	err := sr.Repo.ReActivateService(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}