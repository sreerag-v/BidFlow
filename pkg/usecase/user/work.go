package userUsecase

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
)

type WorkUsecase struct {
	Repo interfaces.WorkRepo
}

func NewWorkUsecase(repo interfaces.WorkRepo) services.WorkUsecase {
	return &WorkUsecase{
		Repo: repo,
	}
}

func (work *WorkUsecase) ListNewOpening(model domain.Work) error {

	//pass to repository
	err := work.Repo.ListNewOpening(model)
	if err != nil {
		return err
	}

	return nil
}