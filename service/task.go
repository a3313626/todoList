package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"

	"github.com/sirupsen/logrus"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	//	Status  int    `json:"status" form:"status"` //0未做 1:已做

}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0未做 1:已做

}

type SearchTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	//Status  int    `json:"status" form:"status"` //0未做 1:已做

	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
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

//显示这个用户所有备忘录
func (service *ListTaskService) ListTask(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0

	if service.PageSize <= 0 {
		service.PageSize = 10
	}

	if service.PageNum <= 0 {
		service.PageNum = 1
	}

	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).First(&tasks, uid).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Order("id desc").Find(&tasks)

	if count == 0 {
		return serializer.Response{
			Status: 500,
			Msg:    "您没有创建备忘录,请先创建",
		}
	}

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

//修改一条备忘录
func (service *UpdateTaskService) UpdateTask(uid uint, taskId string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, taskId).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "没有找到这条备忘录记录,无法修改",
		}
	}

	if task.Uid != uid {
		return serializer.Response{
			Status: 500,
			Msg:    "没有找到这条备忘录记录",
		}
	}

	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	if task.Status == 1 {
		task.EndTime = time.Now().Unix()
	}

	err = model.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "修改失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "",
		Data:   serializer.BuildTask(task),
		//Data:   task,
	}

}

//显示这个用户所有备忘录
func (service *SearchTaskService) SearchTask(uid uint, title string) serializer.Response {
	logrus.Info(title)
	logrus.Info("查看提示")
	var tasks []model.Task
	count := 0

	if service.PageSize <= 0 {
		service.PageSize = 10
	}

	if service.PageNum <= 0 {
		service.PageNum = 1
	}

	model.DB.Model(&model.Task{}).Preload("User").Where("title like ?", "%"+title+"%").Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Order("id desc").Find(&tasks)

	if count == 0 {
		return serializer.Response{
			Status: 500,
			Msg:    "没有找到你想要的备忘录",
		}
	}

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}
