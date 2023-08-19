package types

type Team struct {
	ID          int32  `json:"id"`
	RId         int32  `json:"rId"`
	Name        string `json:"name"`
	CapId       int32  `json:"capId"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type GetTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}

type GetTeamResponse struct {
	Team       Team   `json:"team"`
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}
type GetAllTeamsRequest struct {
	ID       int32  `json:"id"`
	Token    string `json:"token"`
	PageInfo PageInfo
}

type GetAllTeamsResponse struct {
	Teams      []Team `json:"teams"`
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}

type BuildTeamRequest struct {
	TeamName string `json:"teamname"`
}
type BuildTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}

type JoinTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type JoinTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}
type DismissTeamRequest struct {
	Rid   string `json:"rid"`
	Tid   string `json:"tid"`
	Token string `json:"token"`
}
type DismissTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}
type UploadRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type UploadResponse struct {
	Md5        string `json:"md5"`
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}
type QuitTeamRequest struct {
	Rid string `json:"rId"`
	Tid string `json:"tId"`
}
type QuitTeamResponse struct {
	StatusMsg  string `json:"statusMsg"`
	StatusCode int32  `json:"statusCode"`
}
