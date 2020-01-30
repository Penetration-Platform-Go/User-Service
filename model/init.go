package model

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

// User define
type User struct {
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

// Result define
type Result struct {
	IsValid bool
	Value   string
}

// MysqlGrpcClient for connection auth grpc service
var MysqlGrpcClient *grpc.ClientConn

func init() {
	// get user service address
	MYSQLADDRESS := "localhost:8082"
	conn, err := grpc.Dial(MYSQLADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	MysqlGrpcClient = conn
}
