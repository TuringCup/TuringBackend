package types

import "github.com/TuringCup/TuringBackend/repository/db/model"

type GetTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}

type GetTeamResponse struct {
	Team model.Team `json:"team"`
}
