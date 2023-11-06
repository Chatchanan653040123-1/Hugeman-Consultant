package handlers

import (
	"fmt"
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
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Body Parser"})
		return
	}

	response, err := h.backendTestsService.CreateATask(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h backendTestsHandler) ShowSortData(c *gin.Context) {
	sortRequest := services.SortRequest{}
	if err := c.ShouldBindJSON(&sortRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Body Parser"})
		return
	}
	response, err := h.backendTestsService.ShowSortData(sortRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h backendTestsHandler) ShowSearchData(c *gin.Context) {
	searchRequest := services.SearchRequest{}
	if err := c.ShouldBindJSON(&searchRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Body Parser"})
		return
	}
	response, err := h.backendTestsService.ShowSearchData(searchRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h backendTestsHandler) UpdateATask(c *gin.Context) {
	request := services.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Body Parser"})
		fmt.Println("error0")
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		fmt.Println("error1")
		return
	}
	response, err := h.backendTestsService.UpdateATask(request, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		fmt.Println("error2")
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h backendTestsHandler) GetAllOfTasks(c *gin.Context) {
	response, err := h.backendTestsService.GetAllOfTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Task Credentials"})
		return
	}
	c.JSON(http.StatusOK, response)
}
