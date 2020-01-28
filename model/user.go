package model

import (
	"context"
	"log"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
	"google.golang.org/grpc"
)

// InsertUser into mysql
func InsertUser(user *User) bool {
	conn, err := grpc.Dial(MySQLClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	mclient := pb.NewMysqlClient(conn)
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
	conn, err := grpc.Dial(MySQLClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	mclient := pb.NewMysqlClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mclient.UpdateUser(ctx, &pb.UserInformation{
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Email:    user.Email,
		Photo:    user.Photo,
	})

	return result.IsVaild
}
