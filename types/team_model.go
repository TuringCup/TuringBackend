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
	Token string `json:"token"`
}
type BuildTeamResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

type JoinTeamRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type JoinTeamResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}
type DismissTeamRequest struct {
	Rid   string `json:"rid"`
	Tid   string `json:"tid"`
	Token string `json:"token"`
}
type DismissTeamResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}
type UploadRequest struct {
	Rid string `json:"rid"`
	Tid string `json:"tid"`
}
type UploadResponse struct {
	Md5       string `json:"md5"`
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}
type QuitTeamRequest struct {
	Rid string `json:"rId"`
	Tid string `json:"tId"`
}
type QuitTeamResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}
