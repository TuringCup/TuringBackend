package routes

import (
	"net/http"

	Api "github.com/TuringCup/TuringBackend/api"
	"github.com/TuringCup/TuringBackend/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	authed := r.Group("/api")
	authed.Use(middleware.AuthMiddleWare())
	{
		user := authed.Group("/user")
		{
			user.PUT("/:id", Api.UserUpdateHandler())
			user.GET("/:id", Api.UserFindHandler())
		}
		race := authed.Group("/race")
		{
			teams := race.Group(":rid/team")
			{
				team := teams.Group("/:tid")
				{
					team.POST("/join")
					team.POST("/upload")
					team.DELETE("/")
					team.DELETE("/quit")
				}
			}
		}
	}
	api := r.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})

		user := api.Group("/user")
		{
			user.GET("/login", Api.UserLoginHandler())
			user.POST("/register", Api.UserRegisterHandler())
			user.POST("/register/validcode", Api.UserRegisterValidCodeHandler())
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
				}

			}
			race.GET("/", Api.RaceFindAllHandler())
			race.GET("/:rid", Api.RaceFindHandler())
			race.GET("/:rid/progress")
		}
		admin := api.Group("/admin")
		{
			admin.POST("/addrace", Api.RaceAddHandler())
		}
	}
}
