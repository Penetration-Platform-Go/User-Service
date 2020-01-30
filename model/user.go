package model

import (
	"context"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// InsertUser into mysql
func InsertUser(user *User) bool {
	mclient := pb.NewMysqlClient(MysqlGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := mclient.InsertUser(ctx, &pb.UserInformation{
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Email:    user.Email,
		Photo:    user.Photo,
	})

	return result.IsVaild
}

// UpdateUser for mysql
func UpdateUser(user *User) bool {
	mclient := pb.NewMysqlClient(MysqlGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := mclient.UpdateUser(ctx, &pb.UserInformation{
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Email:    user.Email,
		Photo:    user.Photo,
	})

	return result.IsVaild
}
