package main

import (
	"fmt"
	"strings"
	_ "testBackend/docs"
	"testBackend/handlers"
	"testBackend/logs"
	"testBackend/repositories"
	"testBackend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend Test(Golang) API
// @description This is a sample server for Backend Test API.
// @BasePath /api
// @schemes http https
// @host localhost:5000
// @version v1
func main() {
	initConfig()
	backendTestsRepo := repositories.NewBackendTestsRepo(viper.GetString("FILE_PATH"))
	backendTestsService := services.NewBackendTestsService(backendTestsRepo)
	backendTestsHandler := handlers.NewBackendTestsHandler(backendTestsService)

	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow all origins
	config.AllowHeaders = []string{
		"Content-Type", "Content-Length", "Accept-Encoding",
		"X-CSRF-Token", "Authorization", "X-Max",
	}
	router.Use(cors.New(config))
	//swagger
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/create", backendTestsHandler.CreateATask)
	router.POST("/api/showSortData", backendTestsHandler.ShowSortData)
	router.POST("/api/showSearchData", backendTestsHandler.ShowSearchData)
	router.POST("/api/update/:id", backendTestsHandler.UpdateATask)

	router.Run(fmt.Sprintf(":%s", viper.GetString("APP_PORT")))
}

func initConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		logs.Error("Error loading .env file")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
