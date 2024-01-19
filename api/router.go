package api

import (
	"student_api_gateway/api/docs"
	h "student_api_gateway/api/handlers"
	"student_api_gateway/config"
	"student_api_gateway/pkg/logger"
	"student_api_gateway/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @title Service for students
// @version 1.0
// @description This swagger is created to illustrate the body of microservice for students
// @termsOfService http://swagger.io/

// @in header
// @name Komron

// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/"

	handler := h.New(&h.HandlerConfig{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router
	api.POST("/student", handler.CreateStudent)
	api.GET("/student/:id", handler.GetStudent)
	api.PUT("/student", handler.UpdateStudent)
	api.DELETE("/student/:id", handler.DeleteStudent)
	api.POST("/students", handler.ListStudents)

	api.POST("/task", handler.CreateTask)
	api.GET("/task/:id", handler.GetTask)
	api.PUT("/task", handler.UpdateTask)
	api.DELETE("/task/:id", handler.DeleteTask)
	api.POST("/tasks", handler.ListTasks)
	api.GET("/tasks",handler.ListOverdueTasks)
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	return router
}