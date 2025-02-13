package models

import (
	"time"
)

type Rsvp struct{
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id"`
	EventID   uint `json:"event_id"`
	Status    string `json:"status"`
	ResponseDate *time.Time `json:"response_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}