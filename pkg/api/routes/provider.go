package routes

import (
	"github.com/gin-gonic/gin"
	proiderHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func ProviderRoutes(engine *gin.RouterGroup,
	providerHandler *proiderHandler.ProviderHandler,
	profileHandler *proiderHandler.ProfileHandler,
	proworkHandler *proiderHandler.ProWorkHandler) {

	engine.POST("/register", providerHandler.Register)
	engine.POST("/login", providerHandler.Login)

	engine.Use(middleware.ProviderAuthMiddleware)
	{
		profile := engine.Group("/profile")
		{
			profile.GET("", profileHandler.GetDetailsOfProviders)
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

		work := engine.Group("/works")
		{
			leads := work.Group("/leads")
			{
				leads.GET("/list-all", proworkHandler.GetAllLeads)
				leads.GET("/view/:id", proworkHandler.ViewLeads)
			}

			bid := work.Group("bids")
			{
				bid.POST("/place-bid/:id", proworkHandler.PlaceBid)
				bid.PUT("/replace-bid/:id", proworkHandler.ReplaceBidWithNewBid)
				bid.GET("/compare/:id", proworkHandler.GetAllOtherBidsOnTheLeads)
				bid.GET("/accepted-bids",proworkHandler.GetAllAcceptedBids)
			}

			mywork:=work.Group("my-work")
			{
				mywork.GET("/my-works",proworkHandler.GetWorksOfAProvider)
				mywork.GET("/on-going",proworkHandler.GetAllOnGoingWorks)
				mywork.GET("/completed",proworkHandler.GetCompletedWorks)
			}
		}
	}
}
