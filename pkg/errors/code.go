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
	UserNameUsed       = 10005
	EmailUsed          = 10006
	LoginFailed        = 10007

	// ValidCodeError
	ValidCodeError    = 20001
	ValidCodeGenError = 20002

	// Token
	GenerateTokenError = 30001
)
