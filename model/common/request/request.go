package request

type Ids struct {
	Ids []int `json:"ids" form:"ids"`
}

type Id struct {
	Id int `json:"id" form:"id" binding:"required" msg:"id不能为空"`
}
