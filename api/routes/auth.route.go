package routes

import (
	"github.com/gin-gonic/gin"

	"timecassette_api/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login())
		auth.POST("/signup", controllers.Signup())
	}

	// User routes
	user := router.Group("/user")
	{
		user.GET("/search", controllers.SearchEmail())
	}

	// Cassette routes
	cassettes := router.Group("/cassettes")
	{
		cassettes.GET("/", controllers.GetAllCassettesByUser())
		cassettes.POST("/", controllers.CreateCassette())
		cassettes.PUT("/:id", controllers.UpdateCassette())
		cassettes.DELETE("/:id", controllers.DeleteCassette())

		fragments := cassettes.Group("/:cassetteId/fragments")
		{
			fragments.GET("/", controllers.GetAllFragmentsByCassette())
			fragments.POST("/", controllers.CreateFragment())
			fragments.PUT("/:id", controllers.UpdateFragment())
			fragments.DELETE("/:id", controllers.DeleteFragment())

			branches := fragments.Group("/:fragmentId/branches")
			{
				branches.GET("/", controllers.GetAllBranchesByFragment())
				branches.POST("/", controllers.CreateBranch())
				branches.POST("/:id/confirm", controllers.ConfirmBranch())
				branches.PUT("/:id", controllers.UpdateBranch())
				branches.POST("/:id/delete-request", controllers.DeleteBranchRequest())
				branches.POST("/:id/delete-confirm", controllers.DeleteBranchConfirm())

				times := branches.Group("/:branchId/times")
				{
					times.GET("/", controllers.GetAllTimesByBranch())
					times.POST("/", controllers.CreateTime())
					times.PUT("/:id", controllers.UpdateTime())
					times.PUT("/:id/description", controllers.UpdateTimeDescription())
					times.PUT("/:id/start", controllers.UpdateStartTime())
					times.PUT("/:id/end", controllers.UpdateEndTime())
					times.DELETE("/:id", controllers.DeleteTime())
					times.DELETE("/", controllers.DeleteAllTime())
				}
			}
		}
	}
}
