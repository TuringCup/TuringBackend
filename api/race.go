package api

import (
	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/service"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
				StatusCode: errors.InvalidParams,
				StatusMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		resp := service.RaceAdd(ctx, &req)
		ctx.JSON(http.StatusOK, resp)
	}
}

func RaceFindAllHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.GetAllRacesRequest
		pageStr, isGetPage := ctx.GetQuery("page")
		perPageStr, isGetPerPage := ctx.GetQuery("perPage")
		if !isGetPage || !isGetPerPage {
			resp := types.GetAllRacesResponse{
				StatusCode: errors.InvalidParams,
				StatusMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			resp := types.GetAllRacesResponse{
				StatusCode: errors.InvalidParams,
				StatusMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			resp := types.GetAllRacesResponse{
				StatusCode: errors.InvalidParams,
				StatusMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		if page <= 0 || perPage <= 0 {
			resp := types.GetAllRacesResponse{
				StatusCode: errors.InvalidParams,
				StatusMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		req.PageInfo.Page = int32(page)
		req.PageInfo.PerPage = int32(perPage)
		resp := service.RacePage(ctx, &req)
		ctx.JSON(http.StatusOK, resp)
	}
}
