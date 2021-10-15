package response

var Message = map[int]string{
	Success:                 "success",
	Fail:                    "bad request",
	Error:                   "error",
	Unauthorized:            "unauthorized",
	Forbidden:               "forbidden",
}

// GetMsg 根据状态码转换消息体
func GetMsg(code int) string {
	if msg, err := Message[code]; err {
		return msg
	} else {
		return Message[Fail]
	}
}
