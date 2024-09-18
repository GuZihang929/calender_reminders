package system

import "time"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}

type Calender struct {
	Id         int       `json:"id"`
	Type       int       `json:"type"` // 0为邮箱，1为手机，2为全部
	UserId     int       `json:"user_id"`
	Context    string    `json:"context"`
	RemindTime time.Time `json:"remind_time"`
	Status     int       `json:"status"`
}
