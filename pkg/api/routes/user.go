package routes

import (
	"github.com/gin-gonic/gin"
	proiderHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/provider"
	userHandler "github.com/sreerag_v/BidFlow/pkg/api/handler/user"
	"github.com/sreerag_v/BidFlow/pkg/api/middleware"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *userHandler.UserHandler,
	workHandler *userHandler.WorkHandler,
	proworkHandler *proiderHandler.ProWorkHandler,
	ProProfileHandler *proiderHandler.ProfileHandler) {
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
			profile.POST("sent-otp", userHandler.ForgottPassword)
			profile.POST("change-password", userHandler.ChangePassword)
		}

		works := engine.Group("/works")
		{
			works.POST("/add", workHandler.ListNewOpening)
			works.GET("/list", workHandler.GetAllListedWorks)
			works.POST("/image", workHandler.AddImageOfWork)

			works.GET("/on-going", workHandler.ListAllOngoingWorks)
			works.GET("/finished", workHandler.ListAllCompletedWorks)
			works.GET("/work-byid/:id", workHandler.WorkDetailsById)
		}

		proprofile := engine.Group("pro-profile")
		{
			proprofile.GET("", ProProfileHandler.GetProDetails)
		}

		workMGMT := works.Group("work-mgmt")
		{
			workMGMT.GET("/allbids", workHandler.GetAllBids)
			workMGMT.GET("/allacceptedbids",workHandler.GetAllAcceptedBids)
			workMGMT.GET("/bids/:id", proworkHandler.GetAllOtherBidsOnTheLeads)
			workMGMT.PUT("/accept-bid/:id", workHandler.AcceptBid)
			workMGMT.PUT("/assign-work/:id", workHandler.AssignWorkToProvider)
			workMGMT.PUT("/work-completed/:id", workHandler.MakeWorkAsCompleted)
			workMGMT.POST("/rate-work/:id", workHandler.RateWork)
		}
	}

	Payment := engine.Group("/payment")
	{
		Payment.GET("/payment/razorpays/:id", workHandler.RazorPaySent)
		Payment.GET("/payment/success", workHandler.RazorPaySucess)
		// Payment.GET("/payment/successok/:id",workHandler.Success)
	}

}
