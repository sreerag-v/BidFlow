package userRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (usr *UserRepo) CheckPhoneNumberExist(num string) (bool, error) {
	var count int64
	if err := usr.DB.Table("users").Where("phone = ?", num).Count(&count).Error; err != nil {
		return true, err
	}

	// If count is greater than 0, it means a record with the given name exists
	return count > 0, nil
}

func (usr *UserRepo) SignUp(Body models.UserSignup) error {
	err := usr.DB.Exec("INSERT INTO users(name,email,password,phone)VALUES($1,$2,$3,$4)", Body.Name, Body.Email, Body.Password, Body.Phone).Error
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserRepo) CheckUsername(name string) (bool, error) {
	var count int64
	if err := usr.DB.Table("users").Where("name = ?", name).Count(&count).Error; err != nil {
		return true, err
	}

	// If count is greater than 0, it means a record with the given name exists
	return count > 0, nil
}

func (usr *UserRepo) GetUserDetails(name string) (domain.User, error) {
	var model domain.User
	if err := usr.DB.Table("users").Where("name = ?", name).Scan(&model).Error; err != nil {
		return domain.User{}, err
	}

	return model, nil
}

func (usr *UserRepo) CheckUserBlockedOrNot(name string) (domain.User, error) {
	var model domain.User
	if err := usr.DB.Table("users").Where("name = ?", name).Scan(&model).Error; err != nil {
		return domain.User{}, err
	}

	return model, nil
}

func (usr *UserRepo) GetUserDetailsById(id uint) (domain.User, error) {
	var model domain.User
	if err := usr.DB.Table("users").Where("id = ?", id).Scan(&model).Error; err != nil {
		return domain.User{}, err
	}

	return model, nil
}

func (usr *UserRepo) UserProfile(ctx context.Context, id int) ([]models.UserDetails, error) {
	if ctx.Err() != nil {
		return []models.UserDetails{}, errors.New("timeout")
	}

	var user []models.UserDetails
	err := usr.DB.
		Table("users").Where("id = ?", id).Scan(&user).
		Error

	if err != nil {
		return []models.UserDetails{}, fmt.Errorf("Something error")
	}

	return user, nil
}

func (usr *UserRepo) UpdateProfile(id int, body models.UpdateUser) error {
	err := usr.DB.Table("users").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":  body.Name,
			"email": body.Email,
			"phone": body.Phone,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (usr *UserRepo) FindUserByEmail(email string) (domain.User, error) {
	var body domain.User
	err := usr.DB.Table("users").Where("email = ?", email).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (usr *UserRepo) ForgottPassword(body models.Forgott, otp string) error {
	err := usr.DB.Table("users").Where("email = ?", body.Email).Update("otp", otp).Error
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserRepo) ChangePassword(body models.ChangePassword) error {
	err := usr.DB.Table("users").Where("email = ?", body.Email).Update("password", body.Password).Error
	if err != nil {
		return err
	}

	return nil
}
