package response

var Message = map[int]string{
	Success:                 "success",
	Fail:                    "bad request",
	Error:                   "error",
	Unauthorized:            "unauthorized",
	Forbidden:               "forbidden",
	ErrorExist:              "errorExist",
	ErrorNotExist:           "errorNotExist",
	ErrorCreateFail:         "errorCreateFail",
	ErrorUpdateFail:         "errorUpdateFail",
	ErrorDeleteFail:         "errorDeleteFail",
	ErrorSelectFail:         "errorSelectFail",
	ErrorAuthCheckTokenFail: "errorAuthCheckTokenFail",
	ErrorAuthRoleFail:       "errorAuthRoleFail",
}

// GetMsg 根据状态码转换消息体
func GetMsg(code int) string {
	if msg, err := Message[code]; err {
		return msg
	} else {
		return Message[Fail]
	}
}
