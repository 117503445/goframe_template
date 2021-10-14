package model

type RoleApiRequest struct {
	Name string `orm:"name"       json:"name"` //
}
type RoleApiResponse struct {
	Id   uint64 `orm:"id,primary" json:"id"`   //
	Name string `orm:"name"       json:"name"` //
}
