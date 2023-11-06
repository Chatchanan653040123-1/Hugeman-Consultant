package services_test

import (
	"testBackend/repositories"
	"testBackend/services"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockBackendTestsRepository is a mock for the BackendTestsRepository.
type MockBackendTestsRepository struct {
	mock.Mock
}

func (m *MockBackendTestsRepository) CreateATask(task repositories.BackendTests) (*repositories.BackendTests, error) {
	args := m.Called(task)
	return args.Get(0).(*repositories.BackendTests), args.Error(1)
}

func (m *MockBackendTestsRepository) UpdateATask(task repositories.BackendTests, id uuid.UUID) (*repositories.BackendTests, error) {
	args := m.Called(task, id)
	return args.Get(0).(*repositories.BackendTests), args.Error(1)
}

func (m *MockBackendTestsRepository) GetAllOfTasks() ([]repositories.BackendTests, error) {
	args := m.Called()
	return args.Get(0).([]repositories.BackendTests), args.Error(1)
}

func TestCreateATask(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Define the request
	request := services.Request{
		Title:       "Test Title",
		Description: "Test Description",
		Image:       "Test Image",
		Status:      "IN_PROGRESS",
	}

	// Mock the repository's CreateATask function
	mockRepo.On("CreateATask", mock.Anything).Return(&repositories.BackendTests{

		ID:        uuid.New(),
		Title:     request.Title,
		CreatedAt: time.Now(),
	}, nil)

	// Call the service function
	response, err := service.CreateATask(request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestCreateATask_ValidationError(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Define a request with invalid data (triggering a validation error)
	request := services.Request{
		Title:  "", // Empty Title
		Status: "INVALID_STATUS",
	}

	// Call the service function
	response, err := service.CreateATask(request)

	// Assertions
	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestShowSortData(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Mock the repository's GetAllOfTasks function
	mockRepo.On("GetAllOfTasks").Return([]repositories.BackendTests{
		{
			ID:        uuid.New(),
			Title:     "Task A",
			CreatedAt: time.Now(),
			Status:    "IN_PROGRESS",
		},
		{
			ID:        uuid.New(),
			Title:     "Task B",
			CreatedAt: time.Now().Add(-24 * time.Hour),
			Status:    "COMPLETE",
		},
	}, nil)
	SortBy := []string{"Title", "Date", "Status"}
	for _, sortBy := range SortBy {
		// Define the SortRequest
		sortRequest := services.SortRequest{SortBy: sortBy}

		// Call the service function
		response, err := service.ShowSortData(sortRequest)

		// Assertions
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, 2, len(response))
		mockRepo.AssertExpectations(t)
	}
}

func TestShowSortData_ValidationError(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Mock the repository's GetAllOfTasks function
	mockRepo.On("GetAllOfTasks").Return([]repositories.BackendTests{}, nil)

	// Define an invalid SortRequest (triggering a validation error)
	sortRequest := services.SortRequest{SortBy: "InvalidField"}

	// Call the service function
	response, err := service.ShowSortData(sortRequest)

	// Assertions
	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestShowSearchData(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Mock the repository's GetAllOfTasks function
	mockRepo.On("GetAllOfTasks").Return([]repositories.BackendTests{
		{
			ID:          uuid.New(),
			Title:       "Task A",
			Description: "Description A",
			CreatedAt:   time.Now(),
			Status:      "IN_PROGRESS",
		},
		{
			ID:          uuid.New(),
			Title:       "Task B",
			Description: "Description B",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			Status:      "COMPLETE",
		},
	}, nil)

	// Define the SearchRequest
	searchRequest := services.SearchRequest{SearchKeyWord: "Description A"}

	// Call the service function
	response, err := service.ShowSearchData(searchRequest)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, 1, len(response))
	mockRepo.AssertExpectations(t)
}

func TestUpdateATask(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Define a request for updating a task
	request := services.Request{
		Title:       "Updated Title",
		Description: "Updated Description",
		Image:       "Updated Image",
		Status:      "COMPLETE",
	}

	// Mock the repository's UpdateATask function
	taskID := uuid.New()
	mockRepo.On("UpdateATask", mock.Anything, taskID).Return(&repositories.BackendTests{
		ID:        taskID,
		Title:     request.Title,
		CreatedAt: time.Now(),
		Status:    request.Status,
	}, nil)

	// Call the service function
	response, err := service.UpdateATask(request, taskID)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestUpdateATask_ValidationError(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(MockBackendTestsRepository)

	// Initialize the service with the mock repository
	service := services.NewBackendTestsService(mockRepo)

	// Define a request for updating a task with invalid data (triggering a validation error)
	request := services.Request{
		Title:  "", // Empty Title
		Status: "INVALID_STATUS",
	}

	// Call the service function
	response, err := service.UpdateATask(request, uuid.New())

	// Assertions
	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestCreateATask_MaxTitleLength(t *testing.T) {
	mockRepo := new(MockBackendTestsRepository)
	service := services.NewBackendTestsService(mockRepo)

	// Create a request with a title that exceeds the maximum length
	request := services.Request{
		Title:       "ThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharactersThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharactersThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharacters",
		Description: "Test Description",
		Image:       "Test Image",
		Status:      "IN_PROGRESS",
	}

	response, err := service.CreateATask(request)

	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestCreateATask_InvalidStatus(t *testing.T) {
	mockRepo := new(MockBackendTestsRepository)
	service := services.NewBackendTestsService(mockRepo)

	// Create a request with an invalid status
	request := services.Request{
		Title:       "Test Title",
		Description: "Test Description",
		Image:       "Test Image",
		Status:      "INVALID_STATUS",
	}

	response, err := service.CreateATask(request)

	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestShowSearchData_NoMatchingResults(t *testing.T) {
	mockRepo := new(MockBackendTestsRepository)
	service := services.NewBackendTestsService(mockRepo)

	mockRepo.On("GetAllOfTasks").Return([]repositories.BackendTests{
		{
			ID:          uuid.New(),
			Title:       "Task A",
			Description: "Description A",
			CreatedAt:   time.Now(),
			Status:      "IN_PROGRESS",
		},
		{
			ID:          uuid.New(),
			Title:       "Task B",
			Description: "Description B",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			Status:      "COMPLETE",
		},
	}, nil)

	// Try to search with a non-matching keyword
	searchRequest := services.SearchRequest{SearchKeyWord: "NonMatchingKeyword"}

	response, err := service.ShowSearchData(searchRequest)

	require.NoError(t, err)
	require.Empty(t, response)
	mockRepo.AssertExpectations(t)
}

func TestUpdateATask_MaxTitleLength(t *testing.T) {
	mockRepo := new(MockBackendTestsRepository)
	service := services.NewBackendTestsService(mockRepo)

	// Update a task with a title that exceeds the maximum length
	request := services.Request{
		Title:       "ThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharactersThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharactersThisTitleIsTooLongAndExceedsTheMaximumCharacterLimitOfOneHundredCharacters",
		Description: "Test Description",
		Image:       "Test Image",
		Status:      "IN_PROGRESS",
	}
	taskID := uuid.New()

	response, err := service.UpdateATask(request, taskID)

	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestUpdateATask_InvalidStatus(t *testing.T) {
	mockRepo := new(MockBackendTestsRepository)
	service := services.NewBackendTestsService(mockRepo)

	// Update a task with an invalid status
	request := services.Request{
		Title:       "Updated Title",
		Description: "Updated Description",
		Image:       "Updated Image",
		Status:      "INVALID_STATUS",
	}
	taskID := uuid.New()

	response, err := service.UpdateATask(request, taskID)

	require.Error(t, err)
	require.Nil(t, response)
	mockRepo.AssertExpectations(t)
}
