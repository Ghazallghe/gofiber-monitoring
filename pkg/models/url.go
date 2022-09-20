package models

import "github.com/google/uuid"

type Url struct {
	BaseModel
	Url        string `json:"url"`
	Threshold  uint   `json:"threshold"`
	UserId     uuid.UUID
	Statistics []Statistics
}
