package errors

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	ERROR                 = 500
	InvalidParams         = 400

	// UserErrorCode
	RegisterFailed     = 10001
	UserNotExist       = 10002
	UserPasswordWrong  = 10003
	SendValidCodeError = 10004
)
