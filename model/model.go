package model

import "time"

type (
	//Log defines a log
	Log struct {
		ID        uint       `json:"id"`
		Speed     int        `json:"speed"`
		Dir       string     `json:"dir"` //direction
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)

func (Log) TableName() string {
	return "log"
}
