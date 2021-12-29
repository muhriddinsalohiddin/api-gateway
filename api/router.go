package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/muhriddinsalohiddin/api-gateway/api/docs" //swag
	v1 "github.com/muhriddinsalohiddin/api-gateway/api/handlers/v1"
	"github.com/muhriddinsalohiddin/api-gateway/config"
	"github.com/muhriddinsalohiddin/api-gateway/pkg/logger"
	"github.com/muhriddinsalohiddin/api-gateway/services"
)

// Option...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(o Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         o.Logger,
		ServiceManager: o.ServiceManager,
		Cfg:            o.Conf,
	})

	api := router.Group("/v1")
	api.POST("/tasks", handlerV1.CreateTask)
	api.GET("/tasks/:id", handlerV1.GetTask)
	api.GET("/tasks", handlerV1.ListTasks)
	api.PUT("/tasks/:id", handlerV1.UpdateTask)
	api.DELETE("/tasks/:id", handlerV1.DeleteTask)
	api.GET("/taskslist", handlerV1.ListOverdueTask)

	url := ginSwagger.URL("swagger/doc.json") //the url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
