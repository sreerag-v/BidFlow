package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type WorkUsecase interface {
	ListNewOpening(model domain.ReqWork) error
	GetAllListedWorks(id int) ([]models.WorkDetails, error)
	AddImageOfWork(string,int)error
	ListAllCompletedWorks(id int) ([]models.WorkDetails, error)
	ListAllOngoingWorks(id int) ([]models.WorkDetails, error)

	WorkDetailsById(id int) (models.WorkDetails, error)

}
