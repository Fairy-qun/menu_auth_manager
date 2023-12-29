package api

import (
	"github.com/gin-gonic/gin"
	"fairy/service"
	"fairy/utils"
	"fairy/casbin"
	"fmt"
)
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data any `json:"data"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SysApi struct {}

// @Summary 注册用户
// @Description 注册用户数据
// @Tags Sys-系统相关
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /register [post]
func (s SysApi) Register(c *gin.Context) {
	username := c.DefaultPostForm("username","")
	password := c.DefaultPostForm("password","")

	if username == "" || password == "" {
		c.JSON(200,Response{201,"用户名或者密码不能为空",[]interface{}{}})
		return
	}
	users := service.FindUserByName(username)
	if (len(users) != 0) {
		c.JSON(200,Response{201,"用户已经存在,请登录",[]interface{}{}})
		return
	}
	// 注册用户
	// 首先将密码加密处理
	hashPwd,_ := utils.Enbcrypt(password)
	isRegister := service.RegisterUser(username,hashPwd)
	if (!isRegister) {
		c.JSON(200,Response{201,"用户注册失败,请重试",[]interface{}{}})
		return
	}
	// 为注册用户创建角色
	casbin := casbin.InitCasbin()
	casbin.AddRoleForUser(username,"user")
	c.JSON(200,Response{200,"用户注册成功",[]interface{}{}})
}

// @Summary 用户登录
// @Description 用户登录系统
// @Tags Sys-系统相关
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /login [post]
func (s SysApi) Login(c *gin.Context) {
	var user User
	c.ShouldBindJSON(&user)
	username := user.Username
	password := user.Password

	if username == "" || password == "" {
		c.JSON(200,Response{201,"用户名或者密码不能为空",[]interface{}{}})
		return
	}
	users := service.FindUserByName(username)
	if (len(users) == 0) {
		c.JSON(200,Response{201,"用户不存在,请注册",[]interface{}{}})
		return
	}

	// 解密密码
	isEqual := utils.Debcrypt(users[0].Password,password)
	if !isEqual {
		c.JSON(200,Response{201,"密码错误",[]interface{}{}})
		return
	}
	// 生成token
	token, _ := utils.GenerateToken(users[0].Id, users[0].Username)

	// 获取用户所属角色
	casbin := casbin.InitCasbin()
	role,_ := casbin.GetRolesForUser(username)
	menus := service.GetMenuData(role)
	// 构造菜单数据
	userInfo := map[string]interface{}{
		"id":users[0].Id,
		"username":users[0].Username,
		"menus": menus,
	}
	data := map[string]interface{}{
		"userInfo": userInfo,
		"token": token,
	}
	c.JSON(200,Response{200,"登录成功",data})
}




