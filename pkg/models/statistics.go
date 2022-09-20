package models

import (
	"github.com/google/uuid"
)

type Statistics struct {
	Date        string    `json:"date" gorm:"primaryKey"`
	Url         string    `json:"url"`
	UrlId       uuid.UUID `json:"-" gorm:"primaryKey"`
	StatusOk    uint      `json:"2xx"`
	ClientError uint      `json:"4xx"`
	ServerError uint      `json:"5xx"`
}
