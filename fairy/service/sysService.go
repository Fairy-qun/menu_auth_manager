package service

import (
	"fairy/model"
	"fairy/utils"
)

// 查找用户根据姓名
func FindUserByName(username string) []model.SysUser {
	// 查找数据
	var users []model.SysUser
	utils.DB.Where("username=?",username).First(&users)
	return users
}

// 用户注册
func RegisterUser(username string,password string) bool {
	user := model.SysUser{Username:username,Password:password}
	result := utils.DB.Create(&user)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// 获取菜单数据
func GetMenuData(role []string) []*model.SysMenu{
	var menu []model.SysMenu
	if role[0] == "admin" {
		utils.DB.Find(&menu)
	} else {
		utils.DB.Where("role=?",role[0]).Find(&menu)
	}
	menus := ArrayToTree(menu,0)
	return menus
}

// 菜单构造数据
func ArrayToTree(arr []model.SysMenu, pid int) []*model.SysMenu {
	var newArr []*model.SysMenu
	for i := range arr {
		if arr[i].ParentId == pid {
			children := ArrayToTree(arr, arr[i].Id)
			arr[i].Children = children
			newArr = append(newArr, &arr[i])
		}
	}
	return newArr
}