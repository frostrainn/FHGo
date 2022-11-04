package conf

type Host struct {
	Redis
	MySQL
	PgSql
	ETCD
	RocketMQ
	Nacos
}

type Redis struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}

type MySQL struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}

type PgSql struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}

type ETCD struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}

type RocketMQ struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}

type Nacos struct {
	IP       string
	Port     string
	UserName string
	PassWord string
}
