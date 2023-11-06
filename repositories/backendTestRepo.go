package repositories

import (
	"time"

	"github.com/google/uuid"
)

type BackendTests struct {
	ID          uuid.UUID `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"Created At Date Time"`
	Image       string    `json:"Image"`
	Status      string    `json:"Status"`
}
type BackendTestsRepository interface {
	CreateATask(BackendTests) (*BackendTests, error)
	UpdateATask(BackendTests, uuid.UUID) (*BackendTests, error)
	GetAllOfTasks() ([]BackendTests, error)
}
