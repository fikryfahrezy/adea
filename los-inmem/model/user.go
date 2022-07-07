package model

import "time"

type User struct {
	IsOfficer   bool
	Id          string
	Username    string
	Password    string
	CreatedDate time.Time
}
