package adminRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type UserMgmtRepo struct{
	Db *gorm.DB
}

func NewUserMgmtRepo (DB *gorm.DB)interfaces.UserMgmtRepo{
	return &UserMgmtRepo{
		Db: DB,
	}
}


func (mg *UserMgmtRepo)	GetProviders(ctx context.Context,page models.PageNation) ([]models.ProviderDetails, error){
	if ctx.Err() != nil {
		return []models.ProviderDetails{}, errors.New("timeout")
	}

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var provider []models.ProviderDetails
	err := mg.Db.
		Table("providers").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&provider).
		Error

	if err != nil {
		return []models.ProviderDetails{}, fmt.Errorf("failed to get categories from the database")
	}

	return provider, nil
}

func (mg *UserMgmtRepo)	MakeProviderVerified(ctx context.Context, id int) error{
	tx := mg.Db.Begin()
	err := tx.Exec("UPDATE providers SET is_verified = true WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (mg *UserMgmtRepo) RevokeVerification(ctx context.Context, id int) error {
	tx := mg.Db.Begin()
	err := tx.Exec("UPDATE providers SET is_verified = false WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (mg UserMgmtRepo)	GetUsers(ctx context.Context,page models.PageNation) ([]models.UserDetails, error){
	if ctx.Err() != nil {
		return []models.UserDetails{}, errors.New("timeout")
	}

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var user []models.UserDetails
	err := mg.Db.
		Table("providers").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&user).
		Error

	if err != nil {
		return []models.UserDetails{}, fmt.Errorf("failed to get categories from the database")
	}

	return user, nil
}

func (mg *UserMgmtRepo)	BlockUser(ctx context.Context, id int) error{
	tx := mg.Db.Begin()
	err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (mg *UserMgmtRepo) UnBlockUser(ctx context.Context, id int) error {
	tx := mg.Db.Begin()
	err := tx.Exec("UPDATE users SET is_blocked = false WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (mg *UserMgmtRepo) GetAllPendingVerifications(ctx context.Context) ([]models.Verification, error) {
	if ctx.Err() != nil {
		return []models.Verification{}, errors.New("timeout")
	}
	var verifications []models.Verification
	err := mg.Db.Raw("SELECT id,name FROM providers WHERE is_verified = false").Scan(&verifications).Error
	if err != nil {
		return []models.Verification{}, err
	}

	return verifications, nil
}