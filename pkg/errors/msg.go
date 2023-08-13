package errors

var ErrorMsg = map[int]string{
	SUCCESS:               "成功",
	UpdatePasswordSuccess: "修改密码成功",
	ERROR:                 "Fail",
	InvalidParams:         "请求参数错误",
	RegisterFailed:        "注册失败",
}

func GetMsg(code int) string {
	msg, ok := ErrorMsg[code]
	if ok {
		return msg
	}
	return ErrorMsg[ERROR]
}
