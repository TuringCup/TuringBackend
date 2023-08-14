package routes

import (
	"net/http"

	Api "github.com/TuringCup/TuringBackend/api"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})

		user := api.Group("/user")
		{
			user.GET("/login")
			user.POST("/register", Api.UserRegisterHandler())
			user.POST("/register/validcode", Api.UserRegisterValidCodeHandler())
			user.PUT("/:id")
			user.GET("/:id", Api.UserFindHandler())
			user.GET("/refreshtoken")
		}

		race := api.Group("/race")
		{
			teams := race.Group(":rid/team")
			{
				team := teams.Group("/:tid")
				teams.GET("/ping", func(ctx *gin.Context) {
					rid := ctx.Param("rid")
					ctx.JSON(http.StatusOK, rid)
				})

				teams.GET("/")
				teams.POST("/")

				{
					team.GET("/")
					team.POST("/join")
					team.POST("/upload")
					team.DELETE("/")
					team.DELETE("/quit")
				}

			}
			race.GET("/")
			race.GET("/:rid")
			race.GET("/:rid/progress")
		}

	}
}
