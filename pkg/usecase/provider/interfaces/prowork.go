package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProWorkUsecase interface{
	GetAllLeads(int, models.PageNation) ([]models.WorkDetails, error)
	ViewLeads(int,models.PageNation) (models.WorkDetails, error)
	PlaceBid(models.PlaceBid) error
	ReplaceBidWithNewBid(models.PlaceBid) error
	GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error)
	GetWorksOfAProvider(int,models.PageNation) ([]models.WorkDetails, error)
	GetAllOnGoingWorks(int,models.PageNation) ([]models.WorkDetails, error)
	GetCompletedWorks(int,models.PageNation) ([]models.WorkDetails, error)

	GetAllAcceptedBids(int,models.PageNation)([]domain.Bid,error)
}