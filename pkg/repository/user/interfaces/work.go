package interfaces

import "github.com/sreerag_v/BidFlow/pkg/domain"

type WorkRepo interface {
	ListNewOpening(domain.Work) error
}
