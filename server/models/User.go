package models

type User struct {
	Id           int
	Username     string
	PasswordHash string
	Level        int
	Status       int // 0 正常状态， 1删除

	CreateTime int64
}

type Org struct {
	Id          int
	OrgName     string
	Status      int
	LeaderId    int
	ParentOrgId int

	CreateTime int64
}
