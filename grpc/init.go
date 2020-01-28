package grpc

// UserService Struct
type UserService struct {
}

// MySQLClient define
var MySQLClient string

func init() {
	MySQLClient = "localhost:8082"
}
