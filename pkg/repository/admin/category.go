package adminRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) interfaces.CategoryRepo {
	return &CategoryRepo{
		DB: db,
	}
}

func (adm *CategoryRepo) CheckCategory(ctx context.Context, Cat string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	if err := adm.DB.Table("categories").Where("category = ?", Cat).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}
	// If count is greater than 0, it means a record with the given name exists
	return true, nil
}

func (adm *CategoryRepo) CreateCategory(ctx context.Context, Cat string) error {
	tx := adm.DB.Begin()
	err := tx.Exec("INSERT INTO categories(category) VALUES($1)", Cat).Error
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

func (adm *CategoryRepo) ListCatgory(ctx context.Context, page models.PageNation) ([]domain.Category, error) {
	if ctx.Err() != nil {
		return []domain.Category{}, errors.New("timeout")
	}

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var categories []domain.Category
	err := adm.DB.
		Table("categories").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&categories).
		Error

	if err != nil {
		return categories, fmt.Errorf("failed to get categories from the database")
	}

	return categories, nil
}

func (adm *CategoryRepo) CheckCategoryById(ctx context.Context,id int)(bool,error){
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	if err := adm.DB.Table("categories").Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}
	// If count is greater than 0, it means a record with the given name exists
	return true, nil
}

func (adm *CategoryRepo) DeleteCategory(ctx context.Context, id int) error {
	tx := adm.DB.Begin()
	err := tx.Exec("UPDATE categories SET is_deleted = TRUE WHERE id = $1", id).Error
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
