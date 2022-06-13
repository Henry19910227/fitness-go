package entity

import "time"

func NewMockUsers() []*User {
	users := make([]*User, 0)
	user1 := User{
		ID: 10001,
		Nickname: "Henry",
		Sex: "m",
		Account: "henry@gmail.com",
		Email: "henry@gmail.com",
		Height: 176,
		Weight: 70,
		Birthday: time.Now().Format("2006-01-02"),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	user2 := User{
		ID: 10002,
		Nickname: "Jeff",
		Sex: "m",
		Account: "jeff@gmail.com",
		Email: "jeff@gmail.com",
		Height: 172,
		Weight: 65,
		Birthday: time.Now().Format("2006-01-02"),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	users = append(users, &user1)
	users = append(users, &user2)
	return users
}