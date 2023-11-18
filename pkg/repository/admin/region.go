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

type RegionRepo struct{
	DB *gorm.DB
}

func NewRegionRepo(DB *gorm.DB)interfaces.RegionRepo{
	return &RegionRepo{
		DB: DB,
	}
}

func (reg *RegionRepo)	AddNewState(ctx context.Context,body string) error{
	tx := reg.DB.Begin()
	err := tx.Exec("INSERT INTO states(state) VALUES($1)", body).Error
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

func (reg *RegionRepo)	CheckStateExists(ctx context.Context, body string)(bool,error){
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := reg.DB.Raw("SELECT COUNT(*) FROM states WHERE state = $1", body).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (reg *RegionRepo)ListStates(ctx context.Context, page models.PageNation)([]domain.State,error){
	if ctx.Err() != nil {
		return []domain.State{}, errors.New("timeout")
	}

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var state []domain.State
	err := reg.DB.
		Table("states").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&state).
		Error

	if err != nil {
		return state, fmt.Errorf("failed to get categories from the database")
	}

	return state, nil
}

func (reg RegionRepo)	DeleteState(ctx context.Context,id int) error{
	tx := reg.DB.Begin()
	err := tx.Exec("UPDATE states SET is_deleted = TRUE WHERE id = $1", id).Error
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

func (reg *RegionRepo)	CheckIfDistrictAlreadyExists(ctx context.Context, dis string) (bool, error){
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := reg.DB.Raw("SELECT COUNT(*) FROM districts WHERE district = $1", dis).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (reg *RegionRepo)	AddNewDistrict(ctx context.Context,dis models.AddNewDistrict) error{
	tx := reg.DB.Begin()
	err := tx.Exec("INSERT INTO districts(district,state_id) VALUES($1,$2)", dis.District, dis.StateID).Error
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

func (reg *RegionRepo) GetDistrictsFromState(ctx context.Context, id int) ([]domain.District, error) {
	if ctx.Err() != nil {
		return []domain.District{}, errors.New("timeout")
	}
	var districts []domain.District
	err := reg.DB.Raw("SELECT * FROM districts WHERE state_id = $1", id).Scan(&districts).Error
	if err != nil {
		return []domain.District{}, err
	}

	return districts, nil
}

func (reg *RegionRepo) DeleteDistrictFromState(ctx context.Context, id int) error {
	tx := reg.DB.Begin()
	err := tx.Exec("UPDATE districts SET is_deleted = TRUE WHERE id = $1", id).Error
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