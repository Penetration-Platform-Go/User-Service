package controller

import (
	"strconv"
	"time"

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
	username, exist := ctx.Get("username")
	if !exist || username.(string) != ctx.PostForm("username") {
		ctx.Status(403)
		return
	}
	user := model.User{
		Username: username.(string),
		// Password: lib.StringToMd5(ctx.PostForm("password")),
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

// UploadAvatar Api restore file into {webroot}/platform/
func UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.Status(400)
		return
	}
	path := lib.StringToMd5(strconv.FormatInt(time.Now().UnixNano(), 10)) + ".png"
	err = ctx.SaveUploadedFile(file, photoPath+path)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.String(200, path)
}
