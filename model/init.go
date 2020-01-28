package model

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

// MySQLClient define
var MySQLClient string

func init() {
	MySQLClient = "localhost:8082"
}
