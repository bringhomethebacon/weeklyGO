package handler

import (
	"net/http"
	"strings"
	"weekly-report/middleware"

	"github.com/gin-gonic/gin"
)

func checkStatus(ctx *gin.Context) {
	accept := ctx.Request.Header.Get("Accept")

	if strings.HasPrefix(accept, "application/json") {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		return
	}

	ctx.String(http.StatusOK, "ok")
}

func NewHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// health check
	r.GET("/ok", checkStatus)

	r.POST("/login", Login)
	studentAPI := r.Group("/student", middleware.AuthStudentCheck())
	{
		studentAPI.PUT("/update/Student", UpdateStudent)
	}

	teacherAPI := r.Group("/teacher", middleware.AuthTeacherCheck())
	{
		teacherAPI.POST("/create/student", CreateStudent)
		teacherAPI.GET("/member", GetStudents)
	}
	r.POST("/create/teacher", CreateTeacher)
	return r
}
