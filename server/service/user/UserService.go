package user

type UserService struct{}

func (u *UserService) LogOut() {
	// TODO 退出登录，删除Redis里的Token
}

func (u *UserService) GetToken() {
	// TODO 登录时，计算Token并存入Redis
}

func (u *UserService) Authenticate() {
	// TODO 获取前台传参，加盐、加密后与数据库比对
	//

}
