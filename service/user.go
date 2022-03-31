package service

import (
	"todo_list/model"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	var conut int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&conut)

	if conut <= 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在",
		}
	}

	//判断密码是否正确
	if err := user.CheckPassword(service.Password); err == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误,请重新输入",
		}
	}

	// 发个token
	token, err := utils.GenerateToken(user.ID, user.UserName, user.PassWordDigest)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "登录失败,请稍后重试.错误原因" + err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "登录成功",
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}

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
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	//创建用户,写入数据库
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库写入失败,请重试",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}

}
