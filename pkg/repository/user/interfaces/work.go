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
	AssignWorkToProvider(work_id, pro_id int) error
	CheckWorkCommitOrNot(int)(domain.Work,error)
	MakeWorkAsCompleted(id int) error
	RateWork(models.RatingModel, int) error

	GetAllBids(models.PageNation,int)([]models.BidDetails,error)
	GetAllAcceptedBids(models.PageNation,int)([]models.BidDetails,error)

	FindProviderById(int)(domain.Provider,error)
	FindBidExistOrNot(int,int)(domain.Bid,error)
	AcceptBid(int,int)error

	AddAmountInWork(int,int,float64)error
	GetAmountFromWork(int,int)(float64,error)
	DeleteBids(int,int,int)error

	UpdateWorkPaymentField(int,int)error
	RazorPaySucess(Uid int,Oid string,Pid string,Sig string,total string)error

}
