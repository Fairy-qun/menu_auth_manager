package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"os"
	"fmt"
)

// 初始化casbin服务
func InitCasbin() (*casbin.Enforcer){
	path,_ := os.Getwd()
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/fairy", true)
	e, _ := casbin.NewEnforcer(fmt.Sprintf("%v/config/model.conf",path), a)
	fmt.Println(e)
	return e
}

// // 菜单构造数据
// func ArrayToTree(arr []SysMenu, pid int) []*SysMenu {
// 	var newArr []*SysMenu
// 	for i := range arr {
// 		if arr[i].ParentId == pid {
// 			children := ArrayToTree(arr, arr[i].Id)
// 			arr[i].Children = children
// 			newArr = append(newArr, &arr[i])
// 		}
// 	}
// 	return newArr
// }