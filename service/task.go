package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	//	Status  int    `json:"status" form:"status"` //0未做 1:已做

}

type ShowTaskService struct {
}

//新增一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)

	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "操作成功",
	}

}

//展示一条备忘录
func (service *ShowTaskService) Show(uid uint, taskId string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, taskId).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "没有找到这条备忘录记录,请重试",
		}
	}

	if task.Uid != uid {
		return serializer.Response{
			Status: 500,
			Msg:    "没有找到这条备忘录记录",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "",
		Data:   serializer.BuildTask(task),
		//Data:   task,
	}

}
