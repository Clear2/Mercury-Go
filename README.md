```
c.ShouldBind(&form) --接收的格式—>x-www-form-urlencoded 
c.ShouldBindJSON(&json)  --接收的格式—> raw(json) 
id:=c.Query("id")               
page:=c.DefaultQuery("page", "0")               
name:=c.PostForm("name")    
```
#### go mod 更新
go get -u       
go mod tidy

点击这个链接报名，然后登录极客时间app，就可以学习了哦[机智]：https://u.geekbang.org/subject/intro/1000832?key=ecf174408ad6705c4cd1e8ab8d6e05e8

（如显示已过期，直接登录极客时间就可以观看，证明已经报过名了哦~）

Go 进阶训练营课程大纲：http://gk.link/a/10ALo

Go 进阶训练营核心实战项目：http://gk.link/a/10pom