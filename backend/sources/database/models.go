package database

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	AppName     string    `json:"app_name"`
	WindowTitle string    `json:"window_title"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    int64     `json:"duration"` // Duration in seconds
}

type VscodeActivity struct {
	gorm.Model
	Repo      string    `json:"repo"`
	Branch    string    `json:"branch"`
	File      string    `json:"file"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type Summary struct {
	gorm.Model
	Date      time.Time `json:"date"`
	TotalTime int64     `json:"total_time"`
	Details   string    `json:"details"`
}
