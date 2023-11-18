package adminRepo

import (
	"github.com/sreerag_v/BidFlow/pkg/repository/admin/interfaces"
	"gorm.io/gorm"
)

type ServiceRepo struct{
	DB *gorm.DB
}

func NewServiceRepo(DB *gorm.DB)interfaces.ServiceRepo{
	return &ServiceRepo{
		DB: DB,
	}
}