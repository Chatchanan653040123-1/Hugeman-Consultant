package services

import (
	"time"
)

type BackendTests struct {
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"Created At Date Time"`
	Image       string    `json:"Image"`
	Status      string    `json:"Status"`
}
