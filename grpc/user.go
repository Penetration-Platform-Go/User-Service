package grpc

import (
	"context"
	"fmt"

	user "github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
)

type UserService struct {
}

func (u *UserService) Register(ctx context.Context, req *user.UserInformation) (*user.Result, error) {
	fmt.Println(req.Username)
	return nil, nil
}

func (u *UserService) Login(ctx context.Context, req *user.UserLogin) (*user.Result, error) {
	fmt.Println(req.Username)
	return nil, nil
}
