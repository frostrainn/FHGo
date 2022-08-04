package models

type User struct {
	GlobalModel
	Username     string
	PasswordHash string
	Level        int
	Status       int // 0 正常状态， 1删除

}

type Org struct {
	GlobalModel
	OrgName     string
	Status      int
	LeaderId    int
	ParentOrgId int
}
