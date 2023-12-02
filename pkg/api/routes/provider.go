package routes

import (
	"github.com/gin-gonic/gin"
	proiderHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func ProviderRoutes(engine *gin.RouterGroup,
	providerHandler *proiderHandler.ProviderHandler,
	profileHandler *proiderHandler.ProfileHandler) {

	engine.POST("/register", providerHandler.Register)
	engine.POST("/login", providerHandler.Login)

	engine.Use(middleware.ProviderAuthMiddleware)
	{
		profile := engine.Group("/profile")
		{
			profile.GET("",profileHandler.GetDetailsOfProviders)
			services := profile.Group("/service")
			{
				services.POST("/add-service", profileHandler.AddService)
				services.GET("/list-services", profileHandler.GetSelectedServices)
				services.DELETE("/delete-service", profileHandler.DeleteService)
			}

			location := profile.Group("location")
			{
				location.GET("/list-preferredlocations", profileHandler.GetAllPreferredLocations)
				location.POST("/add-preferredlocations", profileHandler.AddPreferredWorkingLocation)
				location.DELETE("/remove-preferredlocations", profileHandler.RemovePreferredLocation)
			}
		}
	}
}
