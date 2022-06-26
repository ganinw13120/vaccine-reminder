package model

import "time"

type User struct {
	UserName   string    `json:"userName"`
	PersonName string    `json:"personName"`
	BirthDate  time.Time `json:"birthDate"`
}

type UserLogResponse struct {
	UserId        string `json:"userId"`
	PersonName    string `json:"personName"`
	Birth         string `json:"birth"`
	Notifications string `json:"notifications"`
}

type UserLog struct {
	UserId       string
	PersonName   string
	Birth        time.Time
	Notification []string
}

type Files struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Month int    `json:"month"`
}
