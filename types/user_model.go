package types

// login
type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ErrorMsg     string `json:"errorMsg"`
	ErrorCode    int    `json:"errorCode"`
}

// register
type RegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Province string `form:"province"`
	City     string `form:"city"`
	School   string `form:"school"`
	SchoolId string `form:"schoolId"`
	Phone    string `form:"phone"`
	// ValidCode string `form:"validcode"`
}

type RegisterResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode int    `json:"errorCode"`
}

type ValidCodeRequest struct {
	Email string `form:"email"`
}

type ValidCodeResponse struct {
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
	ID string `json:"id"`
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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	School      string `json:"school"`
	SchoolId    string `json:"schoolId"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type UpdateUserResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode int    `json:"errorCode"`
}

type UploadFileResponse struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode int    `json:"errorCode"`
}
