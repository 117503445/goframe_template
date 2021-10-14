package model

type UserRoleApiRequest struct {
}
type UserRoleApiResponse struct {
	UserId uint64 `orm:"user_id,primary" json:"userId"` //
	RoleId uint64 `orm:"role_id,primary" json:"roleId"` //
}
