package models

import "github.com/google/uuid"

type Url struct {
	BaseModel
	Url       string `json:"url"`
	Threshold int32  `json:"threshold"`
	UserId    uuid.UUID
}
