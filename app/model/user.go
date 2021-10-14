package model

type UserApiSignUpReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type UserServiceSignUpReq struct {
	Username string
	Password string
}

type UserApiLoginReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type UserServiceLoginReq struct {
	Username string
	Password string
}

type UserApiRequest struct {
	Username string `orm:"username"   json:"username"` //
	Password string `orm:"password"   json:"password"` //
	Avatar   string `orm:"avatar"     json:"avatar"`   //
}
type UserApiResponse struct {
	Id       uint64 `orm:"id,primary" json:"id"`       //
	Username string `orm:"username"   json:"username"` //
	Password string `orm:"password"   json:"password"` //
	Avatar   string `orm:"avatar"     json:"avatar"`   //
}
