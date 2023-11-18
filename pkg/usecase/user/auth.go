package userUsecase

import (
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	helper "github.com/sreerag_v/BidFlow/pkg/helper/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type UserUsecase struct {
	UserRepo interfaces.UserRepo
	helper   helper.Helper
}

func NewUserUsecase(UserRepo interfaces.UserRepo, helper helper.Helper) services.UserUsecase {
	return &UserUsecase{
		UserRepo: UserRepo,
		helper:   helper,
	}
}

func (usr *UserUsecase) SignUp(Body models.UserSignup) error {
	// check the number exist or not
	exists, err := usr.UserRepo.CheckPhoneNumberExist(Body.Phone)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("phone number already exists")
	}
	hashed, err := usr.helper.CreateHashPassword(Body.Password)
	if err != nil {
		return err
	}

	Body.Password = hashed

	if err := usr.UserRepo.SignUp(Body); err != nil {
		return err
	}

	return nil
}

func (usr *UserUsecase) Login(Body models.Login) (string, error) {
	exists, err := usr.UserRepo.CheckUsername(Body.Username)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("check username again")
	}

	details, err := usr.UserRepo.GetUserDetails(Body.Username)
	if err != nil {
		return "", err
	}

	err = usr.helper.CompareHashAndPassword(details.Password, Body.Password)
	if err != nil {
		return "", errors.New("password mismatch")
	}

	token, err := usr.helper.GenerateTokenUser(details)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (usr *UserUsecase) OtpLogin(body domain.User) (domain.User, error) {
	exists, err := usr.UserRepo.CheckPhoneNumberExist(body.Phone)
	if err != nil {
		return domain.User{}, err
	}

	if !exists {
		return domain.User{}, errors.New("Number not Found")
	}

	details, err := usr.UserRepo.GetUserDetails(body.Name)

	if err != nil {
		return domain.User{}, err
	}

	return details, nil
}

func (usr *UserUsecase) GetUserDetails(body domain.User) (domain.User, error) {

	details, err := usr.UserRepo.GetUserDetailsById(uint(body.ID))

	if details.ID == 0 {
		return domain.User{}, errors.New("User Does Not Exist in this id")
	}
	if err != nil {
		return domain.User{}, err
	}

	return details, nil
}

func (usr *UserUsecase) GetJwtToken(body domain.User) (string, error) {
	token, err := usr.helper.GenerateTokenUser(body)
	if err != nil {
		return token, err
	}

	return token, nil
}
