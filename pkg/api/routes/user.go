package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/user"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *userHandler.UserHandler,
	workHandler *userHandler.WorkHandler) {
	engine.POST("/signup", userHandler.SignUp)
	engine.POST(("/login"), userHandler.Login)

	engine.POST("/otp-login", userHandler.OtpLogin)
	engine.POST("/otp-verify", userHandler.LoginOtpVerify)

	engine.Use(middleware.UserAuthMiddleware)
	{
		profile := engine.Group("profile")
		{
			profile.GET("/show", userHandler.UserProfile)
			profile.PATCH("update", userHandler.UpdateProfile)
			profile.POST("sent-otp",userHandler.ForgottPassword)
			profile.POST("change-password",userHandler.ChangePassword)
		}

		works:=engine.Group("/works")
		{
			works.POST("/add",workHandler.ListNewOpening)
			works.GET("/list",workHandler.GetAllListedWorks)
			works.POST("/image",workHandler.AddImageOfWork)

			works.GET("/on-going",workHandler.ListAllOngoingWorks)
			works.GET("/finished",workHandler.ListAllCompletedWorks)
			works.GET("/work-byid/:id",workHandler.WorkDetailsById)
		}
	}
}
