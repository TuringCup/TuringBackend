package types

type GetAllRacesResponse struct {
	ErrorCode int64  `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Races     []Race `json:"races"`
}
type GetRaceRequest struct {
	ID int `json:"id"`
}
type GetRaceResponse struct {
	Race Race `json:"race"`
}

// Race
type Race struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}
