package handler

import (
	"context"
	"net/http"
	pbt "student_api_gateway/genproto/task_service"
	l "student_api_gateway/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create Task
// @Description This is for creating a task
// @Tags Task
// @Accept json
// @Param body body task.CreateTaskReq true "Body request"
// @Produce json
// @Success 201 {object} task.CreateTaskRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /task [post]
func (h *HandlerConfig) CreateTask(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	body := pbt.CreateTaskReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}
	res, err := h.ServiceManager.TaskService().CreateTask(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to create a task", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary Get Task
// @Description This is for getting a task by its ID
// @Tags Task
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} task.GetTaskRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /task/{id} [get]
func (h *HandlerConfig) GetTask(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.TaskService().GetTask(ctx, &pbt.ById{TaskId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to get a student", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Task
// @Description This is for updating a task by its ID
// @Tags Task
// @Accept json
// @Param body body task.UpdateTaskReq true "Request body"
// @Produce json
// @Success 200 {object} task.Success
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /task [put]
func (h *HandlerConfig) UpdateTask(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	body := pbt.UpdateTaskReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	res, err := h.ServiceManager.TaskService().UpdateTask(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to update a student", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete Student
// @Description This is for deleting a task by its ID
// @Tags Task
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} task.Success
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /task/{id} [delete]
func (h *HandlerConfig) DeleteTask(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.TaskService().DeleteTask(ctx, &pbt.ById{TaskId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to delete a student", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary List Tasks
// @Description This is for getting a list of tasks
// @Tags Tasks
// @Accept json
// @Param body body task.ListTasksReq true "Limit and page"
// @Produce json
// @Success 200 {object} task.ListTasksRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /tasks [post]
func (h *HandlerConfig) ListTasks(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	body := pbt.ListTasksReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.TaskService().ListTasks(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to get a list of students", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary List Overdue Tasks
// @Description This is for getting a list of tasks which were not completed in time
// @Tags Tasks
// @Produce json
// @Success 200 {object} task.ListTasksRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /tasks [get]
func (h *HandlerConfig) ListOverdueTasks(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.TaskService().ListOverDue(ctx, &pbt.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to get a list of overdue tasks", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}
