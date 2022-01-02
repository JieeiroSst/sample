package model

import "time"

type Ip struct {
	Id 		int 		`gorm:"primary_key;auto_increment;not_null"`
	Name    string  	`json:"name"`
	Address string  	`json:"address"`
	Time    time.Time   `json:"time"`
}