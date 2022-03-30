package service

import (
	"todo_list/model"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)

	if count >= 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "该用户名已注册,请重试",
		}
	}

	user.UserName = service.UserName
	//密码加密

}
