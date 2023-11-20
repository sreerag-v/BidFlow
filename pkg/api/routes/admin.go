package routes

import (
	"github.com/gin-gonic/gin"
	adminHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/admin"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *adminHandler.AdminHandler,
	categoryHandler *adminHandler.CategoryHandler,
	servicerHandler *adminHandler.ServiceHandler,
	regionHandler *adminHandler.RegionHandler,
	userMgmtHandler *adminHandler.UserMgmtHandler) {

	engine.POST("/signup", adminHandler.AdminSignup)
	engine.POST("/login", adminHandler.AdminLogin)
	engine.DELETE("/delete", adminHandler.DeleteAdmin)

	engine.Use(middleware.AdminAuthMiddleware)
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
				state.POST("/create",regionHandler.AddNewState)
				state.GET("/list",regionHandler.GetStates)
				state.DELETE("/delete",regionHandler.DeleteState)
			}

			district:=region.Group("/district")
			{
				district.POST("/create",regionHandler.AddNewDistrict)
				district.GET("list",regionHandler.GetDistrictsFromState)
				district.DELETE("delete",regionHandler.DeleteDistrictFromState)
			}
		}
	}

}
