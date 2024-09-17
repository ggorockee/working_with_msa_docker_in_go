package domain

import "time"

type Memo struct {
	Id        uint      `json:"id" gorm:"primaryKey:unique;not null"`
	Title     string    `json:"title" mapstructure:"title"`
	Content   string    `json:"content" mapstructure:"content"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

var MemoModel Memo
