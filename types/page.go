package types

type PageRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type PageResponse struct {
	Data       any
	StatusCode int    `json:"errorCode"`
	StatusMsg  string `json:"errorMsg"`
}
