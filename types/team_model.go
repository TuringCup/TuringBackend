package types

type Team struct {
	ID          int64  `json:"id"`
	CapId       string `json:"capId"`
	CreatedTime string `json:"createdTime"`
	Name        string `json:"name"`
	RId         int64  `json:"rId"`
	UpdatedTime string `json:"updatedTime"`
}

type GetTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}

type GetTeamResponse struct {
	Team Team `json:"team"`
}
type GetAllTeamsResponse struct {
	Team []Team `json:"teams"`
}

type BuildTeamRequest struct {
	TeamName string `json:"teamname"`
}
type BuildTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int    `json:"statusCode"`
}

type JoinTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type JoinTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int    `json:"statusCode"`
}
type DismissTeamRequest struct {
	Rid   string `json:"rid"`
	Tid   string `json:"tid"`
	Token string `json:"token"`
}
type DismissTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int    `json:"statusCode"`
}
type UploadRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type UploadResponse struct {
	Md5        string `json:"md5"`
	StatusMsg  string `json:"statusMsg"`
	StatusCode int    `json:"statusCode"`
}
type QuitTeamRequest struct {
	Rid string `json:"rId"`
	Tid string `json:"tId"`
}
type QuitTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int    `json:"statusCode"`
}
