package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type UserMgmtRepo interface{
	GetProviders(context.Context,models.PageNation) ([]models.ProviderDetails, error)
	MakeProviderVerified(ctx context.Context, id int) error
	RevokeVerification(ctx context.Context, id int) error

	GetUsers(context.Context,models.PageNation) ([]models.UserDetails, error)
	BlockUser(ctx context.Context, id int) error
	UnBlockUser(ctx context.Context, id int) error

}