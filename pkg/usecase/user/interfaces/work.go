package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
)

type WorkUsecase interface {
	ListNewOpening(model domain.Work) error
}
