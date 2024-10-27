package entity

import "time"

type User struct {
	Id          int64
	CreatedAt   time.Time
	Login       string
	PassworHash string
}
