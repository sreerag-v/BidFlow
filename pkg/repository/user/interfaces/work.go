package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type WorkRepo interface {
	ListNewOpening(domain.ReqWork) error
	GetAllWorksOfAUser(id int) ([]int, error)
	GetDetailsOfAWork(int) (models.MinWorkDetails, error)
	GetImagesOfAWork(int) ([]string, error)
	FindProviderIdFromWork(int) (int, error)
	FindProviderName(pro_id int) (string, error)

	AddImageOfWork(string,int)error
	GetAllCompletedWorksOfAUser(id int) ([]int, error)
	GetAllOngoingWorksOfAUser(id int) ([]int, error)

}
