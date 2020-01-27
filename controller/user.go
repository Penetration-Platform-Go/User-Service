package controller

import "github.com/gin-gonic/gin"

import "github.com/Penetration-Platform-Go/User-Service/model"

import "fmt"

// CreateUser method for regist api
func CreateUser(ctx *gin.Context) {
	user := model.User{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
		Nickname: ctx.PostForm("nickname"),
		Email:    ctx.PostForm("email"),
		Photo:    ctx.PostForm("photo"),
	}
	result := model.InsertUser(&user)
	fmt.Println(result)
}

// UpdateUser method for regist api
func UpdateUser(ctx *gin.Context) {
	// check jwt

	user := model.User{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
		Nickname: ctx.PostForm("nickname"),
		Email:    ctx.PostForm("email"),
		Photo:    ctx.PostForm("photo"),
	}
	result := model.UpdateUser(&user)
	fmt.Println(result)
}
