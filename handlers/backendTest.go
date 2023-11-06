package handlers

import (
	"net/http"
	"testBackend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type backendTestsHandler struct {
	backendTestsService services.BackendTestsService
}

func NewBackendTestsHandler(backendTestsService services.BackendTestsService) backendTestsHandler {
	return backendTestsHandler{backendTestsService: backendTestsService}
}
func (h backendTestsHandler) CreateATask(c *gin.Context) {
	request := services.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.backendTestsService.CreateATask(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a task"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h backendTestsHandler) ShowSortData(c *gin.Context) {
	sortRequest := services.SortRequest{}
	if err := c.ShouldBindJSON(&sortRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.backendTestsService.ShowSortData(sortRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sorted data"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h backendTestsHandler) ShowSearchData(c *gin.Context) {
	searchRequest := services.SearchRequest{}
	if err := c.ShouldBindJSON(&searchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.backendTestsService.ShowSearchData(searchRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for data"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h backendTestsHandler) UpdateATask(c *gin.Context) {
	request := services.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	response, err := h.backendTestsService.UpdateATask(request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the task"})
		return
	}

	c.JSON(http.StatusOK, response)
}
