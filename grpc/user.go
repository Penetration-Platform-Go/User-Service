package grpc

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
	user "github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
	"google.golang.org/grpc"
)

// UserService Struct
type UserService struct {
}

// GetInformationByUsername Method
func (u *UserService) GetInformationByUsername(ctx context.Context, username *user.Username) (*user.UserInformation, error) {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	mclient := pb.NewMysqlClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := mclient.QueryUserByUsername(ctx, &pb.Username{Username: username.Username})
	if result.Username == "" {
		return &user.UserInformation{}, errors.New("Username not exists")
	}
	return &user.UserInformation{
		Username: result.Username,
		Nickname: result.Nickname,
		Password: result.Password,
		Email:    result.Email,
		Photo:    result.Photo,
	}, nil
}
