package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "eddycjy", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func main() {
	_ = GetUserInfo()
}
