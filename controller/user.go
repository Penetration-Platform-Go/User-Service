package controller

import (
	"github.com/Penetration-Platform-Go/User-Service/lib"
	"github.com/Penetration-Platform-Go/User-Service/model"
	"github.com/gin-gonic/gin"
)

// CreateUser method for regist api
func CreateUser(ctx *gin.Context) {
	user := model.User{
		Username: ctx.PostForm("username"),
		Password: lib.StringToMd5(ctx.PostForm("password")),
		Nickname: ctx.PostForm("nickname"),
		Email:    ctx.PostForm("email"),
		Photo:    ctx.PostForm("photo"),
	}

	if !lib.VerifyUserFormat(&user) {
		ctx.Status(406)
		return
	}

	result := model.InsertUser(&user)

	if result {
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}

// UpdateUser method for regist api
func UpdateUser(ctx *gin.Context) {
	// check jwt

	user := model.User{
		Username: ctx.PostForm("username"),
		Password: lib.StringToMd5(ctx.PostForm("password")),
		Nickname: ctx.PostForm("nickname"),
		Email:    ctx.PostForm("email"),
		Photo:    ctx.PostForm("photo"),
	}

	if !lib.VerifyUserFormat(&user) {
		ctx.Status(406)
		return
	}

	result := model.UpdateUser(&user)
	if result {
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}
