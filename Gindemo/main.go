package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {
	//配置
	router := gin.Default()
	router.StaticFS("/static", http.Dir("./static"))
	router.Delims("{{{","}}}")
	router.LoadHTMLFiles("index.html","users.html","login.html")

	//页面
	router.GET("/users",users)//用户信息界面

	//操作函数
	router.GET("/api/getuser",getuser)
	router.POST("/api/adduser",adduser)
	router.POST("/api/deluser",deluser)
	router.POST("/api/logintest",logintest)

	router.GET("/",login)//登录密码
	router.Use(authMiddleware)
	router.GET("/index",index)//主页面
	//自定义端口
	http.ListenAndServe(":9090",router)
}













































