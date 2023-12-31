package errors

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	ERROR                 = 500
	InvalidParams         = 400
	Forbidden             = 403

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


	//RaceErrorCode
	RaceNotExist = 30002
	RaceNameUsed = 30005

	// Token
	GenerateTokenError = 20003
	TokenTimeout       = 20004

)
