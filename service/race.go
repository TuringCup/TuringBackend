package service

import (
	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RaceFind(ctx *gin.Context, req *types.GetRaceRequest) (resp *types.GetRaceResponse) {
	racedao := dao.NewRaceDao(ctx)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp = &types.GetRaceResponse{
			Race:       types.Race{},
			StatusCode: errors.InvalidParams,
			StatusMsg:  errors.GetMsg(errors.InvalidParams),
		}
		return resp
	}
	race, err := racedao.FindRaceById(id)
	if err != nil {
		resp = &types.GetRaceResponse{
			Race:       types.Race{},
			StatusCode: errors.RaceNotExist,
			StatusMsg:  errors.GetMsg(errors.RaceNotExist),
		}
		return resp
	}
	resp = &types.GetRaceResponse{
		Race: types.Race{
			ID:          race.ID,
			Name:        race.Name,
			CreatedTime: race.CreatedAt.String(),
			UpdatedTime: race.UpdatedAt.String(),
		},
		StatusCode: errors.SUCCESS,
		StatusMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}

func RaceAdd(ctx *gin.Context, req *types.AddRaceRequest) (resp *types.AddRaceResponse) {
	racedao := dao.NewRaceDao(ctx)
	_, exist, err := racedao.ExistOrNotByRaceName(req.Name)
	if err != nil {
		resp = &types.AddRaceResponse{
			StatusCode: errors.ERROR,
			StatusMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	if exist {
		resp = &types.AddRaceResponse{
			StatusCode: errors.RaceNameUsed,
			StatusMsg:  errors.GetMsg(errors.RaceNameUsed),
		}
		return resp
	}
	err = racedao.CreateRace(&model.Race{
		Name: req.Name,
	})
	if err != nil {
		resp = &types.AddRaceResponse{
			StatusCode: errors.ERROR,
			StatusMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	resp = &types.AddRaceResponse{
		StatusCode: errors.SUCCESS,
		StatusMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}

func RacePage(ctx *gin.Context, req *types.GetAllRacesRequest) (resp *types.GetAllRacesResponse) {
	racedao := dao.NewRaceDao(ctx)
	races, err := racedao.FindRaceByPage(req.PageInfo.Page, req.PageInfo.PerPage)
	if err != nil {
		resp := &types.GetAllRacesResponse{
			StatusCode: errors.ERROR,
			StatusMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	var racesInType []types.Race
	for _, race := range races {
		raceTemp := types.Race{
			ID:          race.ID,
			Name:        race.Name,
			CreatedTime: race.CreatedAt.String(),
			UpdatedTime: race.UpdatedAt.String(),
		}
		racesInType = append(racesInType, raceTemp)
	}
	resp = &types.GetAllRacesResponse{
		Races:      racesInType,
		StatusCode: errors.SUCCESS,
		StatusMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}
