package interfaces

import (
	"context"

	"github.com/sreerag_v/BidFlow/pkg/domain"
)

type AdminRepo interface {
	FindAdminByEmail(ctx context.Context,Email string)(int64,error)
	AdminSignup(ctx context.Context ,body domain.Admin)(error)
	GetAdminDetailsByEmail(ctx context.Context,email string)(domain.Admin,error)
	DeleteAdmin(ctx context.Context,id int)error
}