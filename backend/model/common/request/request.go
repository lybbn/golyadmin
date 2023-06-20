package request

//公共请求体

type Ids struct {
	Ids []int `json:"ids" form:"ids"` //ID列表
}

type Id struct {
	Id int `json:"id" form:"id" binding:"required" msg:"id不能为空"` //id
}

type Empty struct{}
