package httpserver

import (
	"log"

	"github.com/gin-gonic/gin"
	adminHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/admin"
	providerHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	userHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/user"
	"github.com/sreerag_v/BidFlow/pkg/api/routes"
	ws "github.com/sreerag_v/BidFlow/pkg/api/chat"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(adminHandler *adminHandler.AdminHandler,
	categoryHandler *adminHandler.CategoryHandler,
	servicerHandler *adminHandler.ServiceHandler,
	regionHandler *adminHandler.RegionHandler,
	userMgmtHAndler *adminHandler.UserMgmtHandler,

	profileHandler *providerHandler.ProfileHandler,
	providerHanlder *providerHandler.ProviderHandler,
	proworkHandler *providerHandler.ProWorkHandler,

	userHandler *userHandler.UserHandler,
	userworkhandler *userHandler.WorkHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("templates/*.html")
	wsHandler := InitHub()

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, servicerHandler, regionHandler, userMgmtHAndler, wsHandler)
	routes.ProviderRoutes(engine.Group("/provider"), providerHanlder, profileHandler, proworkHandler)
	routes.UserRoutes(engine.Group("/user"), userHandler, userworkhandler, proworkHandler, profileHandler)

	return &ServerHttp{engine: engine}
}

func (server *ServerHttp) Start() {
	err := server.engine.Run(":8080")
	if err != nil {
		log.Fatal("unable to Start Server")
	}
}

func InitHub() *ws.Handler {

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	return wsHandler
}
