package userUsecase

import (
	"context"
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
	block,err:=usr.UserRepo.CheckUserBlockedOrNot(Body.Username)
	if err!=nil{
		return "",err
	}
	if block.IsBlocked {
		return "",errors.New("User Blocked by the Admin")
	}
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

func (usr *UserUsecase)	UserProfile(ctx context.Context, id int)([]models.UserDetails,error){
	users, err := usr.UserRepo.UserProfile(ctx, id)
	if err != nil {
		return []models.UserDetails{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.UserDetails{}, errors.New("request timeout")
	}

	return users, nil
}

func (usr *UserUsecase)	UpdateProfile(uid int,body models.UpdateUser)error{
	exist,err:=usr.UserRepo.GetUserDetailsById(uint(uid))
	if err != nil {
		return err
	}

	if exist.ID == 0 {
		return errors.New("something wrong cant find claim id")
	}

	err=usr.UserRepo.UpdateProfile(uid,body)
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserUsecase)	ForgottPassword(body models.Forgott,otp string)error{
	exist,err:=usr.UserRepo.FindUserByEmail(body.Email)
	if err != nil {
		return err
	}
	if exist.Email == "" {
		return errors.New("email does not exist")
	}

	return usr.UserRepo.ForgottPassword(body,otp)
}

func (usr *UserUsecase)	ChangePassword(body models.ChangePassword)error{
	exist,err:=usr.UserRepo.FindUserByEmail(body.Email)
	if err != nil {
		return err
	}
	if exist.Email == "" {
		return errors.New("user not found")
	}
	if body.Password != body.ConfirmPassword {
		return errors.New("password not match")
	}

	if body.Otp != exist.Otp {
		return errors.New("invalid otp")
	}
	hashed, err := usr.helper.CreateHashPassword(body.Password)
	if err != nil {
		return err
	}

	body.Password = hashed

	err=usr.UserRepo.ChangePassword(body)
	if err!=nil{
		return err
	}
	return nil
}
