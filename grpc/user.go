package grpc

import (
	"context"
	"sync"
	"github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
)


type UserService struct {
	mu sync.Mutex
}

func (u *UserService) Register(ctx context.Context, req *user.UserInformation) (*user.Result, error) {
	fmt.Println(req.username)
	return nil
}