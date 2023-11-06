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

// CreateTask godoc
// @Summary Create a task
// @Description ID, Title,Date, and Status fields must be required
// @Description The Title field must not over 100 characters
// @Description The Status field must accept only IN_PROGRESS or COMPLETE status
// @Tags Task
// @Accept  json
// @Produce  json
// @Param request body services.Request{} true "Request"
// @Success 200 {object} services.Response
// @Router /create [post]
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

// ShowSortData godoc
// @Summary Show sorted data
// @Description To sort the data by Title or Date or Status fields
// @Tags Task
// @Accept  json
// @Produce  json
// @Param request body services.SortRequest{} true "SortRequest"
// @Success 200 {object} services.Response
// @Router /showSortData [post]
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

// ShowSearchData godoc
// @Summary Show searched data
// @Description To search the data by Title or Description fields
// @Tags Task
// @Accept  json
// @Produce  json
// @Param request body services.SearchRequest{} true "SearchRequest"
// @Success 200 {object} services.Response
// @Router /showSearchData [post]
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

// UpdateATask godoc
// @Summary Update a task
// @Description To update Title, Description, Date, Image, and Status fields corresponding to the requirements from the CREATE feature
// @Tags Task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param request body services.Request{} true "Request"
// @Success 200 {object} services.Response
// @Router /update/{id} [post]
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
