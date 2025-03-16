package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"time"
)

type User struct {
	ID         int64
	Username   string
	Password   string
	CreateTime int64
	Admin      bool `gorm:"-"`
}

func (User) TableName() string {
	return "users"
}

func InitDB() *gorm.DB {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "root"  //密码
	host := "localhost" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "gorm"    //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return db

}

func Create(db *gorm.DB) {
	user := User{
		Username:   "create",
		Password:   "password",
		CreateTime: time.Now().Unix(),
	}
	db.Create(&user)

	db.Select("username", "password").Create(&user)

	db.Omit("username").Create(&user)

	var users = []User{{Username: "wafer1"}, {Username: "wafer2"}, {Username: "wafer3"}}
	db.Create(&users)

	var userss = []User{{Username: "wafers1"}, {Username: "wafers2"}}

	db.CreateInBatches(&userss, 1)

	//必须先用model指明表
	db.Model(&User{}).Create(map[string]interface{}{
		"Username": "map_create", "Password": "password",
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Username": "maps_create1", "Password": "password"},
		{"Username": "maps_create2", "Password": "password"},
	})

	db.Model(&User{}).Create(map[string]interface{}{
		"username": "wafer_clause",
		"password": clause.Expr{SQL: "md5(?)", Vars: []interface{}{"123456"}},
	})

	db.Exec("insert into users (username,password,create_time) values (?,?,?)", user.Username, user.Password, user.CreateTime)

}
