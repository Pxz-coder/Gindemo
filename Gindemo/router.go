package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录
func login(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{})
}


//主页面
func index(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",gin.H{})
}

//用户操作界面
func users(c *gin.Context){
	c.HTML(http.StatusOK,"users.html",gin.H{})
}

//Api-getuser
func getuser(c *gin.Context){
	c.Writer.Header().Set("Content-Type","application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	db := Connect()
	Users:=user_find(db)
	json.NewEncoder(c.Writer).Encode(Users)
}

//Api-adduser
func adduser(c *gin.Context){
    Id:=c.PostForm("Id")
    Name:=c.PostForm("Name")
	db := Connect()
	user_add(db,Id,Name);
}

//Api-deluser
func deluser(c *gin.Context){
	Id:=c.PostForm("Id")
	db := Connect()
	user_del(db,Id)
}

//Api-logintest
func logintest(c *gin.Context){
	Id:=c.PostForm("Id")
	Password:=c.PostForm("Password")
	db := Connect()
	result:=user_test(db,Id,Password)
	if(result==true){
		token := CreateToken(Id, Password)
		json.NewEncoder(c.Writer).Encode(token)
	} else {
		json.NewEncoder(c.Writer).Encode(404)
	}
}




