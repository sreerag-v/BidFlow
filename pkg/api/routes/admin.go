package routes

import (
	"github.com/gin-gonic/gin"
	ws "github.com/sreerag_v/BidFlow/pkg/api/chat"
	adminHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/admin"
	// "github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *adminHandler.AdminHandler,
	categoryHandler *adminHandler.CategoryHandler,
	servicerHandler *adminHandler.ServiceHandler,
	regionHandler *adminHandler.RegionHandler,
	userMgmtHandler *adminHandler.UserMgmtHandler,
	wsHandler *ws.Handler) {

	engine.POST("/signup", adminHandler.AdminSignup)
	engine.POST("/login", adminHandler.AdminLogin)
	engine.DELETE("/delete", adminHandler.DeleteAdmin)

	// engine.Use(middleware.AdminAuthMiddleware)
	{
		category := engine.Group("category")
		{
			category.POST("/create", categoryHandler.CreateCategory)
			category.GET("list", categoryHandler.ListCatgory)
			category.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		region := engine.Group("/region")
		{
			state := region.Group("/state")
			{
				state.POST("/create", regionHandler.AddNewState)
				state.GET("/list", regionHandler.GetStates)
				state.DELETE("/delete", regionHandler.DeleteState)
			}

			district := region.Group("/district")
			{
				district.POST("/create", regionHandler.AddNewDistrict)
				district.GET("list", regionHandler.GetDistrictsFromState)
				district.DELETE("delete", regionHandler.DeleteDistrictFromState)
			}
		}

		ProManagement := engine.Group("/provider")
		{
			ProManagement.GET("/get-pro", userMgmtHandler.GetProviders)
			ProManagement.PATCH("/verify-pro", userMgmtHandler.MakeProvidersVerified)
			ProManagement.PATCH("/revoke-pro", userMgmtHandler.RevokeVerification)
			ProManagement.GET("/get-pending-verification", userMgmtHandler.GetAllPendingVerifications)
		}

		UserManagement := engine.Group("/user")
		{
			UserManagement.GET("/get-users", userMgmtHandler.GetUsers)
			UserManagement.PATCH("/block-user", userMgmtHandler.BlockUser)
			UserManagement.PATCH("/unblock-user", userMgmtHandler.UnBlockUser)
		}

		Service := engine.Group("/services")
		{
			Service.POST("/create", servicerHandler.AddServiceToCategory)
			Service.GET("/", servicerHandler.GetAllServices)
			Service.GET("/list", servicerHandler.GetServicesInACategory)
			Service.DELETE("/delete", servicerHandler.DeleteService)
			Service.PATCH("/reactivate", servicerHandler.ReActivateService)

		}

		chat := engine.Group("/ws")
		{
			chat.POST("/createRoom", wsHandler.CreateRoom)
			chat.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
			chat.GET("/getRooms", wsHandler.GetRooms)
			chat.GET("/getClients/:roomId", wsHandler.GetClients)
		}
	}
}
