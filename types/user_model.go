package types

// login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ErrorMsg     string `json:"errorMsg"`
	ErrorCode    string `json:"errorCode"`
}

// register
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Province string `json:"province"`
	City     string `json:"city"`
	School   string `json:"school"`
	SchoolId string `json:"schoolId"`
	Phone    string `json:"phone"`
}

type RegisterResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode int    `json:"errorCode"`
}

// refreshtoken
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}
type GetUserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	School      string `json:"school"`
	SchoolId    string `json:"schoolId"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
	ErrorMsg    string `json:"errorMsg"`
	ErrorCode   int    `json:"errorCode"`
}

type UpdateUserRequest struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Phone       *string `json:"phone"`
	Email       string  `json:"email"`
	School      string  `json:"school"`
	SchoolId    string  `json:"schoolId"`
	CreatedTime string  `json:"createdTime"`
	UpdatedTime string  `json:"updatedTime"`
}

type UpdateUserResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}
