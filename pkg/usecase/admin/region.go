package adminUsecase

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type RegionUsecase struct {
	Repo interfaces.RegionRepo
}

func NewRegionUsecase(repo interfaces.RegionRepo) services.RegionUsecase {
	return &RegionUsecase{
		Repo: repo,
	}
}

func (reg *RegionUsecase) AddNewState(ctx context.Context, body string) error {
	exists, err := reg.Repo.CheckStateExists(ctx, body)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("state already exists")
	}

	err = reg.Repo.AddNewState(ctx, body)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (reg *RegionUsecase) ListStates(ctx context.Context, page models.PageNation) ([]domain.State, error) {
	Category, err := reg.Repo.ListStates(ctx, page)
	if err != nil {
		return []domain.State{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.State{}, errors.New("request timeout")
	}
	return Category, nil
}

func (reg *RegionUsecase) DeleteState(ctx context.Context, id int) error {
	exists, err := reg.Repo.CheckStateExistsByid(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("state not exist in this id")
	}
	err = reg.Repo.DeleteState(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (reg *RegionUsecase) AddNewDistrict(ctx context.Context, dis models.AddNewDistrict) error {
	StExist, err := reg.Repo.CheckStateExistsByid(ctx, dis.StateID)
	if err != nil {
		return err
	}
	if StExist {
		return errors.New("There is No State in this id")
	}
	exist, err := reg.Repo.CheckIfDistrictAlreadyExists(ctx, dis.District)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("district already exists")
	}
	//create new category
	err = reg.Repo.AddNewDistrict(ctx, dis)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (reg *RegionUsecase) GetDistrictsFromState(ctx context.Context, id int) ([]domain.District, error) {

	districts, err := reg.Repo.GetDistrictsFromState(ctx, id)
	if err != nil {
		return []domain.District{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.District{}, errors.New("request timeout")
	}

	return districts, nil
}

func (reg *RegionUsecase) DeleteDistrictFromState(ctx context.Context, id int) error {
	exist, err := reg.Repo.CheckIfDistrictExistByid(ctx, id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("district not found in this id")

	}
	err = reg.Repo.DeleteDistrictFromState(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}
