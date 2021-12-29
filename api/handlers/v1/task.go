package v1

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/gin-gonic/gin"
	pb "github.com/muhriddinsalohiddin/api-gateway/genproto"
	"github.com/muhriddinsalohiddin/api-gateway/pkg/logger"
	_ "github.com/muhriddinsalohiddin/api-gateway/api/handlers/models"
	"github.com/muhriddinsalohiddin/api-gateway/pkg/utils"
)

// CreateTask ...
// @Summary CreateTask
// @Description This API for creating a new task
// @Tags task
// @Accept  json
// @Produce  json
// @Param Task request body models.Task true "taskCreateRequest"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks/ [post]
func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body        pb.Task
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.TaskService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to create task")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetTask ...
// @Summary GetTask
// @Description This API for getting task detail
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks/{id} [get]
func (h *handlerV1) GetTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.TaskService().Get(ctx, &pb.ByIdReq{Id: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get task", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListTasks ...
// @Summary ListTasks
// @Description This API for getting list of tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListTasks
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks [get]
func (h *handlerV1) ListTasks(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.TaskService().List(ctx, &pb.ListReq{
		Limit: params.Limit,
		Page:  params.Page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list task", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTask ...
// @Summary UpdateTask
// @Description This API for updating task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Task request body models.UpdateTask true "taskUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks/{id} [put]
func (h *handlerV1) UpdateTask(c *gin.Context) {
	var (
		body        pb.Task
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failet to bind json", logger.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.TaskService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failet to update task", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteTask ...
// @Summary DeleteTask
// @Description This API for deleting task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks/{id} [delete]
func (h *handlerV1) DeleteTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.TaskService().Delete(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete task", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListOverdueTask ...
// @Summary ListOverdueTask
// @Description This API for getting list of tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Param Task request body models.ListOverdue true "taskListOverdueRequest"
// @Success 200 {object} models.ListTasks
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/taskslist [get]
func (h *handlerV1) ListOverdueTask(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var (
		body        pb.ListOverReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("should to bind json", logger.Error(err))
		return
	}
	body.Limit = params.Limit
	body.Page = params.Page
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.TaskService().ListOverdue(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list overude", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}
