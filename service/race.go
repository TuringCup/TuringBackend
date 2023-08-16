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
			Race:      types.Race{},
			ErrorCode: errors.InvalidParams,
			ErrorMsg:  errors.GetMsg(errors.InvalidParams),
		}
		return resp
	}
	race, err := racedao.FindRaceById(id)
	if err != nil {
		resp = &types.GetRaceResponse{
			Race:      types.Race{},
			ErrorCode: errors.RaceNotExist,
			ErrorMsg:  errors.GetMsg(errors.RaceNotExist),
		}
		return resp
	}
	resp = &types.GetRaceResponse{
		Race: types.Race{
			ID:          int(race.ID),
			Name:        race.Name,
			CreatedTime: race.CreatedAt.String(),
			UpdatedTime: race.UpdatedAt.String(),
		},
		ErrorCode: errors.SUCCESS,
		ErrorMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}

func RaceAdd(ctx *gin.Context, req *types.AddRaceRequest) (resp *types.AddRaceResponse) {
	racedao := dao.NewRaceDao(ctx)
	_, exist, err := racedao.ExistOrNotByRaceName(req.Name)
	if err != nil {
		resp = &types.AddRaceResponse{
			ErrorCode: errors.ERROR,
			ErrorMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	if exist {
		resp = &types.AddRaceResponse{
			ErrorCode: errors.RaceNameUsed,
			ErrorMsg:  errors.GetMsg(errors.RaceNameUsed),
		}
		return resp
	}
	err = racedao.CreateRace(&model.Race{
		Name: req.Name,
	})
	if err != nil {
		resp = &types.AddRaceResponse{
			ErrorCode: errors.ERROR,
			ErrorMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	resp = &types.AddRaceResponse{
		ErrorCode: errors.SUCCESS,
		ErrorMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}
