## Golang 学习项目
学习用途

## 目录结构
~~~
├─config                  配置文件
│
├─controller                  控制器
│
├─database               	  数据库连接层
│
├─middlewares        中间件
│
├─router        路由
│
├─models         模型
│
├─service             业务逻辑
│
├─main.go
├─README.md             README 文件
~~~

## QUICK START
	govendor sync
	go run main.go

## 访问
	ip:8000/login

## TODO

- [x] 引入jwt
- [x] 接入数据库
- [x] 接入Hprose