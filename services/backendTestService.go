package services

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID          uuid.UUID `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"Created At Date Time"`
	Image       string    `json:"Image"`
	Status      string    `json:"Status"`
}
type Request struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	Status      string `json:"Status"`
}
type SortRequest struct {
	SortBy string `json:"SortBy"`
}
type SearchRequest struct {
	SearchKeyWord string `json:"SearchKeyWord"`
}
type BackendTestsService interface {
	CreateATask(Request) (*Response, error)
	ShowSortData(SortRequest) ([]Response, error)
	ShowSearchData(SearchRequest) ([]Response, error)
	UpdateATask(Request, uuid.UUID) (*Response, error)
}
