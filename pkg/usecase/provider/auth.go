package providerUsecase

import (
	"errors"

	helper "github.com/sreerag_v/BidFlow/pkg/helper/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProviderUsecase struct {
	repo   interfaces.ProviderRepo
	helper helper.Helper
}

func NewProviderUsecase(Repo interfaces.ProviderRepo, help helper.Helper) services.ProviderUsecase {
	return &ProviderUsecase{
		repo:   Repo,
		helper: help,
	}
}

func (pro *ProviderUsecase) Register(model models.ProviderRegister) error {
	// check the phone number already exist or not
	exists, err := pro.repo.CheckPhoneNumber(model.Phone)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("phone number already exists")
	}

	// check the password matches
	if model.Password != model.RePassword {
		return errors.New("password mismatch")
	}

	hashed, err := pro.helper.CreateHashPassword(model.Password)
	if err != nil {
		return err
	}

	model.Password = hashed

	id, err := pro.repo.Register(model)
	if err != nil {
		return err
	}

	// upload file to aws bucket S3
	filename, err := pro.helper.UploadToS3(model.Document)
	if err != nil {
		return err
	}

	// save that file into database
	err = pro.repo.UploadDoc(id, filename)
	if err != nil {
		return err
	}
	return nil
}

func (pro *ProviderUsecase) Login(Body models.Login) (string, error) {
	// check the usename exist or not
	exists, err := pro.repo.CheckProExistOrNot(Body.Username)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("Username Not Found!!!")
	}

	// fetch the deatils of the user
	details, err := pro.repo.GetProDetails(Body.Username)
	if err != nil {
		return "", err
	}

	// check the passwords
	err = pro.helper.CompareHashAndPassword(details.Password, Body.Password)
	if err != nil {
		return "", errors.New("password mismatch")
	}

	// check verification
	if !details.IsVerified {
		return "", errors.New("your request is under validation")
	}

	token, err := pro.helper.GenerateTokenProvider(details)
	if err != nil {
		return token, err
	}

	return token, nil
}
