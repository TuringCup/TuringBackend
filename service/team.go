package service

import (
	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func TeamFind(ctx *gin.Context, req *types.GetTeamRequest) (resp *types.GetTeamResponse) {
	teamdao := dao.NewTeamDao(ctx)
	tid, err := strconv.Atoi(req.Tid)
	if err != nil {
		resp = &types.GetTeamResponse{
			Team:       types.Team{},
			StatusCode: errors.InvalidParams,
			StatusMsg:  errors.GetMsg(errors.InvalidParams),
		}
		return resp
	}
	rid, err := strconv.Atoi(req.Rid)
	if err != nil {
		resp = &types.GetTeamResponse{
			Team:       types.Team{},
			StatusCode: errors.InvalidParams,
			StatusMsg:  errors.GetMsg(errors.InvalidParams),
		}
		return resp
	}
	team, err := teamdao.FindTeamById(tid, rid)
	if err != nil {
		resp = &types.GetTeamResponse{
			Team:       types.Team{},
			StatusCode: errors.TeamNotExist,
			StatusMsg:  errors.GetMsg(errors.TeamNotExist),
		}
		return resp
	}
	resp = &types.GetTeamResponse{
		Team: types.Team{
			ID:          team.ID,
			RId:         team.Rid,
			CapId:       team.CapId,
			Name:        team.Name,
			CreatedTime: team.CreatedAt.String(),
			UpdatedTime: team.UpdatedAt.String(),
		},
		StatusCode: errors.SUCCESS,
		StatusMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}

func TeamPage(ctx *gin.Context, req *types.GetAllTeamsRequest) (resp *types.GetAllTeamsResponse) {
	teamdao := dao.NewTeamDao(ctx)
	teams, err := teamdao.FindTeamByPage(req.PageInfo.Page, req.PageInfo.PerPage)
	if err != nil {
		resp := &types.GetAllTeamsResponse{
			StatusCode: errors.ERROR,
			StatusMsg:  errors.GetMsg(errors.ERROR),
		}
		return resp
	}
	var teamsInTypes []types.Team
	for _, team := range teams {
		teamTemp := types.Team{
			ID:          team.ID,
			RId:         team.Rid,
			Name:        team.Name,
			CapId:       team.CapId,
			CreatedTime: team.CreatedAt.String(),
			UpdatedTime: team.UpdatedAt.String(),
		}
		teamsInTypes = append(teamsInTypes, teamTemp)
	}
	resp = &types.GetAllTeamsResponse{
		Teams:      teamsInTypes,
		StatusCode: errors.SUCCESS,
		StatusMsg:  errors.GetMsg(errors.SUCCESS),
	}
	return resp
}
