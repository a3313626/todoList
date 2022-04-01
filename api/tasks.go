package api

import (
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTaskService service.CreateTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&createTaskService); err == nil {
		res := createTaskService.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTaskService service.ShowTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&showTaskService); err == nil {
		res := showTaskService.Show(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}

func ListTask(c *gin.Context) {
	var service service.ListTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListTask(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}

func UpdateTask(c *gin.Context) {
	var service service.UpdateTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateTask(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}

func SearchTask(c *gin.Context) {
	var service service.SearchTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchTask(claim.Id, c.PostForm("title"))
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}
