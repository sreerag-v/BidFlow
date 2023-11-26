package adminUsecase

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type UserMgmtUsecase struct {
	Repo        interfaces.UserMgmtRepo
	ServiceRepo interfaces.ServiceRepo
}

func NewUserMgmtUsecase(repo interfaces.UserMgmtRepo, ServiceRepo interfaces.ServiceRepo) services.UserMgmtUsecase {
	return &UserMgmtUsecase{
		Repo:        repo,
		ServiceRepo: ServiceRepo,
	}
}

func (mg *UserMgmtUsecase) GetProviders(ctx context.Context, page models.PageNation) ([]models.ProviderDetails, error) {

	users, err := mg.Repo.GetProviders(ctx, page)
	if err != nil {
		return []models.ProviderDetails{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.ProviderDetails{}, errors.New("request timeout")
	}

	return users, nil
}

func (mg *UserMgmtUsecase) MakeProviderVerified(ctx context.Context, id int) error {
	err := mg.Repo.MakeProviderVerified(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (mg *UserMgmtUsecase) RevokeVerification(ctx context.Context, id int) error {
	err := mg.Repo.RevokeVerification(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (mg *UserMgmtUsecase) GetUsers(ctx context.Context, page models.PageNation) ([]models.UserDetails, error) {
	users, err := mg.Repo.GetUsers(ctx, page)
	if err != nil {
		return []models.UserDetails{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.UserDetails{}, errors.New("request timeout")
	}

	return users, nil
}

func (mg *UserMgmtUsecase) BlockUser(ctx context.Context, id int) error {
	exists, err := mg.Repo.CheckUserExistOrNot(ctx, id)
	if err != nil {
		return err
	}
	if exists.ID != 0 {
		err = mg.Repo.BlockUser(ctx, id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("user does not exist")
		return res
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (mg *UserMgmtUsecase) UnBlockUser(ctx context.Context, id int) error {

	exists, err := mg.Repo.CheckUserExistOrNot(ctx, id)
	if err != nil {
		return err
	}
	if exists.ID != 0 {
		err := mg.Repo.UnBlockUser(ctx, id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("user does not exist")
		return res
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (mg *UserMgmtUsecase) GetAllPendingVerifications(ctx context.Context,page models.PageNation) ([]models.Verification, error) {
	verification, err := mg.Repo.GetAllPendingVerifications(ctx,page)
	if err != nil {
		return []models.Verification{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.Verification{}, errors.New("request timeout")
	}

	return verification, nil
}
