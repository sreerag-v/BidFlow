package adminUsecase

import (
	"context"
	"errors"
	"sync"

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

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	errCh := make(chan error, 2)
	wg.Add(1)

	// first go routine for finding the user exist or not 
	go func() {
		defer wg.Done()
		exist, err := adm.repository.FindAdminByEmail(ctx, body.Email)
		if err != nil {
			errCh <- err
			return
		}
		// check the user exist or not
		if exist != 0 {
			errCh <- errors.New("admin email  already exists")
		}
	}()

	// second go routine for hashing the password
	wg.Add(1)
	go func(){
		defer wg.Done()
		hash, err := adm.helper.CreateHashPassword(body.Password)
		if err != nil {
			errCh<- err
			return
		}
		body.Password = hash
	}()

	// Wait for both goroutines to finish
	go func() {
		wg.Wait()
		close(errCh)
	}()


	// Check for errors from the goroutines
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	
	// give the previlage of admin
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

func (adm *adminUsecase) DeleteAdmin(ctx context.Context, id int) error {
	err := adm.repository.DeleteAdmin(ctx, id)
	if err != nil || ctx.Err() != nil {
		return err
	}

	return nil
}
