package models

import "time"

//Memo model for memo
type Memo struct {
	ID      int        `gorm:"column:id"`
	Title   string     `gorm:"column:title"`
	Body    string     `gorm:"column:body"`
	Created time.Time  `gorm:"column:created_at"`
}

//TableName target table
func (c Memo) TableName() string {
	return "memos"
}