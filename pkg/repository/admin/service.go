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

type ServiceRepo struct {
	DB *gorm.DB
}

func NewServiceRepo(DB *gorm.DB) interfaces.ServiceRepo {
	return &ServiceRepo{
		DB: DB,
	}
}

func (sr *ServiceRepo) CheckIfServiceAlreadyExists(ctx context.Context, service string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := sr.DB.Raw("SELECT COUNT(*) FROM professions WHERE profession = $1", service).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (sr *ServiceRepo) AddServicesToACategory(ctx context.Context, service models.AddServicesToACategory) error {
	tx := sr.DB.Begin()
	err := tx.Exec("INSERT INTO professions(profession,category_id) VALUES($1,$2)", service.ServiceName, service.CategoryID).Error
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

func (sr *ServiceRepo) GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error) {
	if ctx.Err() != nil {
		return []domain.Profession{}, errors.New("timeout")
	}
	var services []domain.Profession
	err := sr.DB.Raw("SELECT * FROM professions WHERE category_id = $1", id).Scan(&services).Error
	if err != nil {
		return []domain.Profession{}, err
	}

	return services, nil
}

func (sr *ServiceRepo) GetAllServices(ctx context.Context, page models.PageNation) ([]domain.Profession, error) {
	if ctx.Err() != nil {
		return []domain.Profession{}, errors.New("timeout")
	}

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var service []domain.Profession
	err := sr.DB.
		Table("professions").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&service).
		Error

	if err != nil {
		return service, fmt.Errorf("failed to get categories from the database")
	}

	return service, nil
}

func (sr *ServiceRepo) DeleteService(ctx context.Context, id int) error {
	tx := sr.DB.Begin()
	err := tx.Exec("UPDATE professions SET is_deleted = TRUE WHERE id = $1", id).Error
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

func (sr *ServiceRepo) ReActivateService(ctx context.Context, id int) error {
	tx := sr.DB.Begin()
	err := tx.Exec("UPDATE professions SET is_deleted = False WHERE id = $1", id).Error
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
