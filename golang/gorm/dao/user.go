package dao

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID         int64
	Username   string
	Password   string
	CreateTime int64
	Admin      bool `gorm:"-"`
}

// static table name
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

// dynamic table name
func UserTable(user *User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if user.Admin {
			return tx.Table("admin_users")
		}

		return tx.Table("users")
	}
}

func Save(user *User) {
	user.Admin = true

	//err := DB.Create(user).Error

	//err := DB.Scopes(UserTable(user)).Create(user).Error

	//err := DB.Table("users").Create(user).Error

	//err := DB.Select("username", "password").Create(&user).Error

	err := GetDB().Omit("username").Create(&user).Error

	if err != nil {
		log.Println("insert fail : ", err)
	}
}

func GetById(id int64) User {
	var user User
	err := GetDB().Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}

func UpdateById(id int64) {
	err := GetDB().Model(&User{}).Where("id=?", id).
		Update("username", "refaw").Error
	if err != nil {
		log.Println("update users  fail : ", err)
	}
}

func DeleteById(id int64) {
	err := GetDB().Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}
