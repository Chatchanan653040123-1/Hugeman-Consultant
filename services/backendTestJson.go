package services

import (
	"encoding/base64"
	"sort"
	"testBackend/errs"
	"testBackend/repositories"
	"time"

	"github.com/google/uuid"
)

type backendTestsService struct {
	backendTestsRepo repositories.BackendTestsRepository
}

func NewBackendTestsService(backendTestsRepo repositories.BackendTestsRepository) backendTestsService {
	return backendTestsService{backendTestsRepo: backendTestsRepo}
}

// The TODO application can CREATE a task with the following requirements
// The ID, Title,Date, and Status fields must be required
// The ID field must be UUID format
// The Title field must not over 100 characters
// The Date field must be RFC3399 with Timezone format
// The Status field must accept only IN_PROGRESS or COMPLETE status
// The Image field must be Based64 Encode format
func (s backendTestsService) CreateATask(req Request) (*Response, error) {
	//validation
	if len(req.Title) == 0 || len(req.Title) > 100 {
		return nil, errs.NewValidationError("Title must be between 1 and 100 characters")
	}
	if req.Status != "IN_PROGRESS" && req.Status != "COMPLETE" {
		return nil, errs.NewValidationError("Status must be either 'IN_PROGRESS' or 'COMPLETE'")
	}
	//time parse to RFC3339
	t, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	//image string parse to base64
	image := base64.StdEncoding.EncodeToString([]byte(req.Image))
	task := repositories.BackendTests{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   t,
		Image:       image,
		Status:      req.Status,
	}
	createdTask, err := s.backendTestsRepo.CreateATask(task)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	taskResponse := Response{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		CreatedAt:   createdTask.CreatedAt,
		Image:       createdTask.Image,
		Status:      createdTask.Status,
	}
	return &taskResponse, nil
}

// Can sort the data by Title or Date or Status fields
func (s backendTestsService) ShowSortData(req SortRequest) ([]Response, error) {
	tasks, err := s.backendTestsRepo.GetAllOfTasks()
	if err != nil {
		return nil, err
	}

	var lessFunc func(i, j int) bool

	switch req.SortBy {
	case "Title":
		lessFunc = func(i, j int) bool {
			return tasks[i].Title < tasks[j].Title
		}
	case "Date":
		lessFunc = func(i, j int) bool {
			return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
		}
	case "Status":
		lessFunc = func(i, j int) bool {
			return tasks[i].Status < tasks[j].Status
		}
	default:
		return nil, errs.NewValidationError("SortBy must be either 'Title', 'Date' or 'Status'")
	}
	sort.Slice(tasks, func(i, j int) bool {
		return lessFunc(i, j)
	})

	var tasksResponse []Response
	for _, task := range tasks {
		taskResponse := Response{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			Image:       task.Image,
			Status:      task.Status,
		}
		tasksResponse = append(tasksResponse, taskResponse)
	}

	return tasksResponse, nil
}

// Can search the data by Title or Description fields
func (s backendTestsService) ShowSearchData(req SearchRequest) ([]Response, error) {
	tasks, err := s.backendTestsRepo.GetAllOfTasks()
	if err != nil {
		return nil, err
	}
	var tasksResponse []Response
	for _, task := range tasks {
		if task.Title == req.SearchKeyWord || task.Description == req.SearchKeyWord {
			taskResponse := Response{
				ID:          task.ID,
				Title:       task.Title,
				Description: task.Description,
				CreatedAt:   task.CreatedAt,
				Image:       task.Image,
				Status:      task.Status,
			}
			tasksResponse = append(tasksResponse, taskResponse)
		}
	}
	return tasksResponse, nil
}
func (s backendTestsService) UpdateATask(req Request, id uuid.UUID) (*Response, error) {
	//validation
	if len(req.Title) == 0 || len(req.Title) > 100 {
		return nil, errs.NewValidationError("Title must be between 1 and 100 characters")
	}
	if req.Status != "IN_PROGRESS" && req.Status != "COMPLETE" {
		return nil, errs.NewValidationError("Status must be either 'IN_PROGRESS' or 'COMPLETE'")
	}
	//time parse to RFC3339
	t, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	//image string parse to base64
	image := base64.StdEncoding.EncodeToString([]byte(req.Image))
	task := repositories.BackendTests{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   t,
		Image:       image,
		Status:      req.Status,
	}
	updatedTask, err := s.backendTestsRepo.UpdateATask(task, id)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	taskResponse := Response{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		CreatedAt:   updatedTask.CreatedAt,
		Image:       updatedTask.Image,
		Status:      updatedTask.Status,
	}
	return &taskResponse, nil
}

// Extra
func (s backendTestsService) GetAllOfTasks() ([]Response, error) {
	tasks, err := s.backendTestsRepo.GetAllOfTasks()
	if err != nil {
		return nil, err
	}
	var tasksResponse []Response
	for _, task := range tasks {
		taskResponse := Response{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			Image:       task.Image,
			Status:      task.Status,
		}
		tasksResponse = append(tasksResponse, taskResponse)
	}
	return tasksResponse, nil
}
