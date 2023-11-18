package routes

import (
	"github.com/gin-gonic/gin"
	proiderHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func ProviderRoutes(engine *gin.RouterGroup,
	providerHandler *proiderHandler.ProviderHandler) {

	engine.POST("/register", providerHandler.Register)
	engine.POST("/login", providerHandler.Login)

	engine.Use(middleware.ProviderAuthMiddleware)
	{
		
	}
}
