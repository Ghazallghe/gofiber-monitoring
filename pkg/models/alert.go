package models

import "github.com/google/uuid"

type Alert struct {
	BaseModel
	ErrorCounts uint `json:"errorCounts"`
	UrlId       uuid.UUID
}
