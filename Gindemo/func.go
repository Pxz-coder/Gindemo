package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//错误处理
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//验证用户
func user_test(db *sql.DB,Id string,Password string) bool {
	rows, err := db.Query("select * from loginMessage where Id=" + "'" + Id + "'" + " and Password =" + "'" + Password + "'")
	checkErr(err)
	var id string
	var password string
	for rows.Next() {
		err = rows.Scan(&id, &password)
		checkErr(err)
	}
	db.Close()
	if (id == Id && password == Password) {
		return true
	} else {
		return false
	}
}


//查找用户
func user_find(db *sql.DB)[]User{
	rows, err := db.Query("select * from users")
	checkErr(err)
	results:=[]User{}
	for rows.Next() {
		single:=User{}
		var Id string
		var Name string
		err = rows.Scan(&Id, &Name)
		single.Id=Id
		single.Name=Name
		results=append(results,single)
		checkErr(err)
	}
	db.Close()
	return results
}

//添加用户
func user_add(db *sql.DB,Id string,Name string){
	stmt, err := db.Prepare("insert into users (Id, Name) values($1, $2)")
	if err != nil {
		fmt.Printf("Insert ERR:, %v", err)
	}
	_, err = stmt.Exec(Id,Name)
	if err != nil {
		fmt.Printf("Insert ERR:, %v", err)
	}
	db.Close()
}

//删除用户
func user_del(db *sql.DB,Id string){
	stmt, err := db.Prepare("delete from users where Id=$1")
	if err != nil {
		fmt.Printf("Update ERR, %v", err)
	}
	_, err = stmt.Exec(Id)
	if err != nil {
		fmt.Printf("Update ERR, %v", err)
	}
	db.Close()
}


//解析Token
func parseToken(s string) (*jwt.Token, error) {
	fn := func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	}
	return jwt.Parse(s, fn)
}


//创建token
func CreateToken(Id, Password string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["Id"] = Id
	claims["Password"] = Password
	token.Claims = claims
	if tokenString, err := token.SignedString([]byte(SecretKey)); err == nil {
		return tokenString
	} else {
		return ""
	}
}

//验证token中间件
func authMiddleware(c *gin.Context) {
	tokenMessage:= c.Query("token")
	tokenMessage = tokenMessage[1 : len(tokenMessage)-1]
	token, err := parseToken(tokenMessage)
	if err != nil {
		c.JSON(200, "token verify error")}
	//如果验证结果是false直接返回token错误我 如果成功则继续下一个handler
	if token.Valid {
		c.Next()
	} else {
		c.Abort()
	}
}




