### 环境准备
```
go
go get -u github.com/gin-gonic/gin
go get -u github.com/go-ini/ini
go get -u github.com/unknwon/com
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/astaxie/beego/validation
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
```


### 实现功能
1. 基础的CRUD
```
详见 gorm官方文档
```
2. 跨域设置
3. 分页设置
4. 原生sql查询与修改更新
5. 异常统一处理
```
还能进一步抽象。
目前新增tag路由的返回，会比较优雅一点。
```
6. 查询条件的筛选
```
模糊查询，和其它复杂的查询，需要进行改造
```
7. 简单jwt接口权限
8. swagger

### article 和 tag 有差异
```
tag 新增标签方法，使用的是body传参，直接传入Tag实体类，根据tag的model,进行参数校验，完成tag的添加
article 使用的是param传参，完成参数校验之后，再新建实例，完成新增操作。
目前来看，新增使用Body传参比较合适，代码量少很多。
```

### 参考项目地址：

https://eddycjy.gitbook.io/golang/
Gin搭建Blog API's 实战项目
https://juejin.cn/post/6863765115456454664
gin参数校验之validator库
https://swaggo.github.io/swaggo.io/declarative_comments_format/
swagger配置文档
https://gorm.io/zh_CN/docs/models.html
gorm 官方文档

### 打包与部署
linux
```
1.打包
set GOARCH=amd64
set GOOS=linux
go build main.go
会得到一个main的文件
2.部署
添加权限
chmod 777 main
守护进程运行
setsid ./main
```
windows
```
go build main.go
会得到一个main的exe文件
管理员权限运行
```
本工程编译速度2s,大小34mb

### 实际开发可以使用模块化设计api
https://github.com/flipped-aurora/gin-vue-admin  
基于gin+vue搭建的后台管理系统框架  
https://github.com/xinliangnote/go-gin-api  
基于 Gin 进行模块化设计的 API 框架，封装了常用功能，使用简单，致力于进行快速的业务研发。