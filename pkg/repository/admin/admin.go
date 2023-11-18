package adminRepo

import (
	"context"
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) interfaces.AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (adm *AdminRepo) FindAdminByEmail(ctx context.Context, email string) (int64, error) {
	var count int64
	err := adm.DB.Raw("SELECT COUNT(*) FROM admins WHERE email = $1", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (adm *AdminRepo) AdminSignup(ctx context.Context, body domain.Admin) error {
	tx := adm.DB.Begin()
	//insert new admin to admin table
	err := tx.Exec("INSERT INTO admins (name,email,password,previlege) VALUES($1,$2,$3,$4)", body.Name, body.Email, body.Password, body.Previlege).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (adm *AdminRepo) GetAdminDetailsByEmail(ctx context.Context, email string) (domain.Admin, error) {
	var model domain.Admin
	err := adm.DB.Table("admins").Where("email = ?", email).Scan(&model).Error
	if err != nil {
		return domain.Admin{}, err
	}
	if ctx.Err() != nil {
		return domain.Admin{}, errors.New("timeout")
	}

	return model, nil
}

func (adm *AdminRepo) DeleteAdmin(ctx context.Context, id int) error {
	tx := adm.DB.Begin()
	err := tx.Table("admins").Where("id = ?", id).Delete(id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if ctx.Err() != nil {
		tx.Rollback()
		return errors.New("request timeout")
	}

	tx.Commit()
	return nil
}
