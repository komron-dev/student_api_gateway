package handler

import (
	"context"
	"net/http"
	pbs "student_api_gateway/genproto/student_service"
	l "student_api_gateway/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create Student
// @Description This is for creating a student
// @Tags Student
// @Accept json
// @Param body body student.CreateStudentReq true "body"
// @Produce json
// @Success 200 {object} student.CreateStudentRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /student [post]
func (h *HandlerConfig) CreateStudent(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	body := pbs.CreateStudentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}
	res, err := h.ServiceManager.StudentService().CreateStudent(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to create a student", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary Get Student
// @Description This is for getting a student by his/her ID
// @Tags Student
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} student.GetStudentRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /student/{id} [get]
func (h *HandlerConfig) GetStudent(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.StudentService().GetStudent(ctx, &pbs.ById{StudentId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to get a student", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Student
// @Description This is for updating a student by his/her ID
// @Tags Student
// @Accept json
// @Param body body student.UpdateStudentReq true "Body"
// @Produce json
// @Success 200 {object} student.Success
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /student [put]
func (h *HandlerConfig) UpdateStudent(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	body := pbs.UpdateStudentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	res, err := h.ServiceManager.StudentService().UpdateStudent(ctx, &body)
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
// @Description This is for deleting a student by his/her ID
// @Tags Student
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} student.Success 
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /student/{id} [delete]
func (h *HandlerConfig) DeleteStudent(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Cfg.ContextTimeout))
	defer cancel()

	res, err := h.ServiceManager.StudentService().DeleteStudent(ctx, &pbs.ById{StudentId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to delete a student", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary List Students
// @Description This is for getting a list of students
// @Tags Students
// @Accept json
// @Param body body student.ListStudentsReq true "Limit and page"
// @Produce json
// @Success 200 {object} student.ListStudentsRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /students [post]
func (h *HandlerConfig) ListStudents(c *gin.Context) {
	var jsMarshal protojson.MarshalOptions
	jsMarshal.UseProtoNames = true

	body := pbs.ListStudentsReq{}
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

	res, err := h.ServiceManager.StudentService().ListStudents(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to get a list of students", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}
