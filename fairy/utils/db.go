package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB
var err error
func init() {
	path,_ := os.Getwd()
	config := viper.New()

	// 设置读取的相关信息
	config.AddConfigPath(fmt.Sprintf("%v/config",path))
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig();err != nil {
		fmt.Println("读取config文件失败:",err)
	}

	// 连接数据库
	host := config.Get("database.host")
	port := config.Get("database.port")
	username := config.Get("database.username")
	password := config.Get("database.password")
	database := config.Get("database.database")
	fmt.Println(database)
	
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",username,password,host,port,database)

	DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败-%v", err)
		return
	}
}