package handler

import (
	"fmt"
	"net/http"

	"weekly-report/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	role := c.Query("role")
	if id == "" || password == "" || role == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "学号、密码和角色不能为空",
		})
		return
	}
	password = model.GetMd5(password)
	var (
		err      error
		userRole string
		username string
	)
	if role == "student" {
		student := model.Student{}
		err = model.DB.Where("id = ? AND password = ? ", id, password).First(&student).Error
		userRole = student.Role
		username = student.StudentName
	} else {
		teacher := model.Teacher{}
		err = model.DB.Where("id = ? AND password = ? ", id, password).First(&teacher).Error
		fmt.Printf("teacher: %v\n", teacher)
		userRole = teacher.Role
		username = teacher.TeacherName
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取用户信息错误:" + err.Error(),
		})
		return
	}

	token, err := model.GenerateToken(userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "GenerateToken Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"token":    token,
		"role":     userRole,
	})
}

func CreateTeacher(c *gin.Context) {
	payload := model.Teacher{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("解析教师信息失败: %v", err),
		})
		return
	}
	payload.Password = "123456"
	// 判断用户名是否已存在
	var cnt int64
	err := model.DB.Where("id = ?", payload.ID).Model(new(model.Teacher)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "查询教师信息失败:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "教师已完成注册",
		})
		return
	}

	data := &model.Teacher{
		ID:          payload.ID,
		Password:    model.GetMd5(payload.Password),
		TeacherName: payload.TeacherName,
		Role:        payload.Role,
	}
	err = model.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "添加教师失败:" + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func CreateStudent(c *gin.Context) {
	payload := model.Student{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("解析学生信息失败: %v", err),
		})
		return
	}
	payload.Password = model.GetMd5("123456")
	// 判断用户名是否已存在
	var cnt int64
	err := model.DB.Where("id = ?", payload.ID).Model(new(model.Student)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "查询学生信息失败:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "该学生已被注册",
		})
		return
	}

	err = model.DB.Create(payload).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "添加学生失败:" + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func UpdateStudent(c *gin.Context) {
	name := c.Query("Studentname")
	password := c.Query("password")
	// 判断用户名是否已存在
	var cnt int64
	err := model.DB.Where("name = ?", name).Model(new(model.Student)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Student Error:" + err.Error(),
		})
		return
	}
	if cnt != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "该用户不存在",
		})
		return
	}

	data := &model.Student{
		Password: model.GetMd5(password),
	}
	err = model.DB.Where("name = ?", name).Updates(data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Crete Student failed:" + err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetStudents(c *gin.Context) {
	Students := []model.Student{}
	err := model.DB.Where("is_admin = ?", 0).Find(&Students).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "get weeklys failed:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"weeklys": Students,
	})
}
