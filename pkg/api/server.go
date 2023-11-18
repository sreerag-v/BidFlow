package httpserver

import (
	"log"

	"github.com/gin-gonic/gin"
	adminHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/admin"
	providerHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	userHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/user"

	"github.com/sreerag_v/BidFlow/pkg/api/routes"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(adminHandler *adminHandler.AdminHandler,
	categoryHandler *adminHandler.CategoryHandler,
	servicerHandler *adminHandler.ServiceHandler,
	regionHandler *adminHandler.RegionHandler,

	providerHanlder *providerHandler.ProviderHandler,
	userHandler *userHandler.UserHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler,categoryHandler,servicerHandler,regionHandler)
	routes.ProviderRoutes(engine.Group("/provider"), providerHanlder)
	routes.UserRoutes(engine.Group("/user"), userHandler)

	return &ServerHttp{engine: engine}
}

func (server *ServerHttp) Start() {
	err := server.engine.Run(":8080")
	if err != nil {
		log.Fatal("unable to Start Server")
	}
}
