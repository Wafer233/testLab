package dao

import (
	"fmt"
	"gorm.io/gorm"
)

type GoodsAPI struct {
	Id    int
	Title string
	//Price      float64
	//Stock      int
	//Type       int
	//CreateTime time.Time
}

// scope
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Scope() {
	db := GetDB()
	var goods []Goods
	db.Scopes(Paginate(2, 3)).Find(&goods)
	fmt.Println("goods", goods)
}

func AdvanceQuery() {
	var goodsAPI GoodsAPI
	db := GetDB()
	db.Model(&Goods{}).Where("id = ?", 3).Find(&goodsAPI)
	fmt.Println("\ngoodsAPI", goodsAPI)
}
