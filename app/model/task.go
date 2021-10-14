package model

type TaskApiResponse struct {
	Id    uint64 `orm:"id,primary" json:"id"`    //
	Title string `orm:"title"      json:"Title"` //
	Done  bool   `orm:"done"       json:"Done"`  //
}

type TaskApiRequest struct {
	Title string `orm:"title"      json:"Title"` //
	Done  bool   `orm:"done"       json:"Done"`  //
}
