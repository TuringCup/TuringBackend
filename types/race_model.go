package types

type GetAllRacesResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Races     []Race `json:"races"`
}
type GetRaceRequest struct {
	ID string `json:"id"`
}
type GetRaceResponse struct {
	Race      Race   `json:"race"`
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

// Race
type Race struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type AddRaceRequest struct {
	Name string `json:"name" form:"name"`
}
type AddRaceResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}
