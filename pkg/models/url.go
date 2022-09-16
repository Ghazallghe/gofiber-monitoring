package models

import "github.com/google/uuid"

type Url struct {
	BaseModel
	Url       string `json:"url"`
	Thershold int32  `json:"thershold"`
	UserId    uuid.UUID
}
