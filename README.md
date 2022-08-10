[toc]

# FHGo

究极缝合怪工程

计划实现各种常用功能

目前已实现：
Zap日志库
LumberJack日志分割存储
Gorm连接MySQL
Go-redis连接Redis
Gin路由组

## Features

算法实现

权限管理

支持MarkDown的博客

支付/聚合支付


## TODO LIST

- [ ] Gin- JWT中间件、Casbin中间件
- [ ] 登录
- [ ] 登陆采用账号密码和token模式。 存放于redis,以cookie的形式传给前端 从前端获取.
- [ ] 验证码
- [ ] 鉴权+权限管理


## 技术栈

 - [X] Redis
 - [X] MySQL
 - [ ] PostgresQL
 - [ ] NSQ
 - [ ] Kafka
 - [ ] 

## 三方库

- [X] Gin/ 高性能网络框架
- [X] Zap/ 日志库
- [X] LumberJack/ 日志分割存储、备份
- [X] gorm/ 数据库操作
- [X] go-redis/ Redis
- [ ] ants/ 高性能Goroutine池
- [ ] go-jwt/ Json Web Token签发
- [ ] casbin/ 访问控制
- [ ] go-micro/ 微服务框架
- [ ] protobuf/ 序列化
- [x] gRPC/ RPC服务
- [ ] Docker/ 容器化部署

 

