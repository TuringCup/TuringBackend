package errors

var ErrorMsg = map[int]string{
	SUCCESS:               "成功",
	UpdatePasswordSuccess: "修改密码成功",
	ERROR:                 "Fail",
	InvalidParams:         "请求参数错误",
	RegisterFailed:        "注册失败",
	UserNotExist:          "用户不存在",
	UserPasswordWrong:     "用户密码错误",
	SendValidCodeError:    "发送验证码失败",
	ValidCodeError:        "发送验证码失败",
	UserNameUsed:          "用户名已经被使用",
	EmailUsed:             "邮箱已经被注册",
	ValidCodeGenError:     "验证码生成错误",
	LoginFailed:           "登录失败",
}

func GetMsg(code int) string {
	msg, ok := ErrorMsg[code]
	if ok {
		return msg
	}
	return ErrorMsg[ERROR]
}
