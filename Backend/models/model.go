package models

import "time"

type Participant struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Name         string    `json:"name" binding:"required"`
	Phone        int       `json:"phone" binding:"required"`
	Age          int       `json:"age" binding:"gte=18,lte=65"`
	Batch        string    `json:"batch" binding:"required"`
	Payment      bool      `json:"payment"`
	EnrollDate   time.Time `json:"enrollDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}
