package types

type GetAllRacesRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type GetAllRacesResponse struct {
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg"`
	Races      []Race `json:"races"`
}
type GetRaceRequest struct {
	ID string `json:"id"`
}
type GetRaceResponse struct {
	Race       Race   `json:"race"`
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg"`
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
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg"`
}
