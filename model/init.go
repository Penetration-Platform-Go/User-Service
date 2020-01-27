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
	isValid bool
	value   string
}

func init() {

}
