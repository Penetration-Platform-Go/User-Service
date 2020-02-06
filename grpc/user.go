package grpc

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
	user "github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
	"google.golang.org/grpc"
)

// GetInformationByUsername Method
func (u *UserService) GetInformationByUsername(ctx context.Context, username *user.Username) (*user.UserInformation, error) {
	conn, err := grpc.Dial(MySQLClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	mclient := pb.NewMysqlClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := mclient.QueryUsers(ctx, &pb.Condition{
		Value: []*pb.Value{
			{Key: "username", Value: username.Username},
		},
	})
	if err != nil {
		return nil, err
	}
	var result user.UserInformation
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		result.Username = feature.Username
		result.Password = feature.Password
		result.Email = feature.Email
		result.Photo = feature.Photo
		result.Nickname = feature.Nickname
	}
	if result.Username == "" {
		return &user.UserInformation{}, errors.New("Username not exists")
	}
	return &result, nil
}
