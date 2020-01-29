package lib

import (
	"crypto/md5"
	"fmt"
	"regexp"

	"github.com/Penetration-Platform-Go/User-Service/model"
)

// VerifyUserFormat check UserInformation format
func VerifyUserFormat(user *model.User) bool {
	if verifyUsernameFormat(user.Username) && verifyEmailFormat(user.Email) &&
		verifyNicknameFormat(user.Nickname) && verifyPhotoFormat(user.Photo) {
		return true
	}
	return false
}

// verifyEmailFormat
func verifyEmailFormat(email string) bool {

	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// verifyUsernameFormat
func verifyUsernameFormat(username string) bool {
	pattern := `^[a-z0-9_-]{6,20}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

// verifyNicknameFormat
func verifyNicknameFormat(nickname string) bool {
	pattern := `^[a-z0-9_-]{3,16}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(nickname)
}

// verifyPhotoFormat
func verifyPhotoFormat(photo string) bool {
	pattern := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(photo)
}

// StringToMd5 Transfer string to Md5
func StringToMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
