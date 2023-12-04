//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	httpserver "github.com/sreerag_v/BidFlow/pkg/api"
	adminHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/admin"
	userHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/user"
	"github.com/sreerag_v/BidFlow/pkg/config"
	"github.com/sreerag_v/BidFlow/pkg/db"
	"github.com/sreerag_v/BidFlow/pkg/helper"
	adminRepo "github.com/sreerag_v/BidFlow/pkg/repository/admin"
	userRepo "github.com/sreerag_v/BidFlow/pkg/repository/user"
	adminUsecase "github.com/sreerag_v/BidFlow/pkg/usecase/admin"
	userUsecase "github.com/sreerag_v/BidFlow/pkg/usecase/user"

	provoderHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	provoderRepo "github.com/sreerag_v/BidFlow/pkg/repository/provider"
	provoderUsecase "github.com/sreerag_v/BidFlow/pkg/usecase/provider"
)

func InitializeAPI(cfg config.Config) (*httpserver.ServerHttp, error) {
	wire.Build(db.ConnectDatabase,
		adminRepo.NewAdminRepo,
		adminUsecase.NewAdminUsecase,
		adminHandler.NewAdminHandler,

		adminHandler.NewCategoryHandler,
		adminUsecase.NewCategoryRepo,
		adminRepo.NewCategoryRepo,

		adminHandler.NewRegionHandler,
		adminUsecase.NewRegionUsecase,
		adminRepo.NewRegionRepo,

		adminHandler.NewUserMgmtHandler,
		adminUsecase.NewUserMgmtUsecase,
		adminRepo.NewUserMgmtRepo,

		adminHandler.NewServiceHandler,
		adminUsecase.NewServiceUsecase,
		adminRepo.NewServiceRepo,

		httpserver.NewServerHttp,
		helper.NewHelper,

		provoderRepo.NewProviderRepo,
		provoderUsecase.NewProviderUsecase,
		provoderHandler.NewProviderHandler,

		provoderRepo.NewProfileRepo,
		provoderUsecase.NewProfileUsecase,
		provoderHandler.NewProfileHandler,

		userRepo.NewUserRepo,
		userUsecase.NewUserUsecase,
		userHandler.NewUserHandler,

		userRepo.NewWorkRepo,
		userUsecase.NewWorkUsecase,
		userHandler.NewWorkHandler,
	)

	return &httpserver.ServerHttp{}, nil
}
