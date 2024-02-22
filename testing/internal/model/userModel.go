package model

type User struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	UserAge  int64  `json:"userAge"`
}
