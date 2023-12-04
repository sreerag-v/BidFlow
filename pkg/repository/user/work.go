package userRepo

import (

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	"gorm.io/gorm"
)

type WorkRepo struct {
	DB *gorm.DB
}

func NewWorkRepo(db *gorm.DB) interfaces.WorkRepo {
	return &WorkRepo{
		DB: db,
	}
}

func (work *WorkRepo) ListNewOpening(model domain.Work) error {

	err := work.DB.Exec("INSERT INTO works(street,district_id,state_id,target_profession_id,user_id)VALUES($1,$2,$3,$4,$5)", model.Street, model.DistrictID, model.StateID, model.TargetProfessionID, model.UserID).Error
	if err != nil {
		return err
	}

	return nil
}


