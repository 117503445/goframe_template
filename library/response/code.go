package response

const (
	Success      = 0     // 请求成功
	Fail         = 40000 // 请求失败
	Error        = 50000 // 服务器内部错误
	Unauthorized = 40001 // 身份未授权
	Forbidden    = 40003 // 路由未授权
)
