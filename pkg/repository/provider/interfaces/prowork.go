package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProWorkRepo interface{
	GetLeadByServiceAndLocation(service, location int) ([]int, error)
	GetDetailsOfAWork(int,models.PageNation) (models.MinWorkDetails, error)
	GetImagesOfAWork(int) ([]string, error)
	CheckIfAlreadyBidded(work_id, pro_id int) (bool, error)
	PlaceBid(model models.PlaceBid) error
	ReplaceBidWithNewBid(model models.PlaceBid) error
	GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error)
	FindProviderName(pro_id int) (string, error)
	GetAllWorksOfAProvider(pro_id int) ([]int, error)
	GetCommittedWorksOfAProvider(pro_id int) ([]int, error)
	GetCompletedWorksOfAProvider(pro_id int) ([]int, error)

	GetAllAcceptedBids(int,models.PageNation)([]domain.Bid,error)
}