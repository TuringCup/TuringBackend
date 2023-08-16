package api

import (
	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/service"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RaceFindHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.GetRaceRequest
		id := ctx.Param("rid")
		req.ID = id
		resp := service.RaceFind(ctx, &req)
		ctx.JSON(http.StatusOK, resp)
	}

}

func RaceAddHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddRaceRequest
		if err := ctx.Bind(&req); err != nil {
			resp := types.AddRaceResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		resp := service.RaceAdd(ctx, &req)
		ctx.JSON(http.StatusOK, resp)
	}
}
