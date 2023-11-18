package adminUsecase

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	helper "github.com/sreerag_v/BidFlow/pkg/helper/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type adminUsecase struct {
	repository interfaces.AdminRepo
	helper     helper.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepo, helper helper.Helper) services.AdminUsecase {
	return &adminUsecase{
		repository: repo,
		helper:     helper,
	}
}

func (adm *adminUsecase) AdminSignup(ctx context.Context, body domain.Admin) error {
	if err := ctx.Err(); err != nil {
		return errors.New("Resuest time Out")
	}

	// chekc the admin exist or not
	exist, err := adm.repository.FindAdminByEmail(ctx, body.Email)

	if err != nil {
		return err
	}

	if exist != 0 {
		return errors.New("admin email  already exists")
	}

	// hash the password
	hash, err := adm.helper.CreateHashPassword(body.Password)

	if err != nil {
		return err
	}

	// give the previlage of admin
	body.Password = hash
	body.Previlege = "admin"

	// create new admin

	if err := adm.repository.AdminSignup(ctx, body); err != nil {
		return err
	}

	return nil
}

func (adm *adminUsecase) AdminLogin(ctx context.Context, Body models.AdminLogin) (string, error) {
	if ctx.Err() != nil {
		return "", errors.New("request time out")
	}

	//check the email already exist or not

	admDetails, err := adm.repository.GetAdminDetailsByEmail(ctx, Body.Email)
	if err != nil {
		return "", err
	}

	//  compare the password
	err = adm.helper.CompareHashAndPassword(admDetails.Password, Body.Password)
	if err != nil {
		return "", err
	}

	var AdminResponse models.AdminDetailsResponse

	AdminResponse.ID = int(admDetails.ID)
	AdminResponse.Email = admDetails.Email
	AdminResponse.Name = admDetails.Name
	AdminResponse.Previlege = admDetails.Previlege

	token, err := adm.helper.GenerateTokenAdmin(AdminResponse)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (adm *adminUsecase) DeleteAdmin(ctx context.Context,id int)error{
	err := adm.repository.DeleteAdmin(ctx, id)
	if err != nil || ctx.Err() != nil {
		return err
	}

	return nil
}
