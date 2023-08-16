package types

type PageRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type PageResponse struct {
	Data      any
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}
