package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Goods struct {
	Id         int
	Title      string
	Price      float64
	Stock      int
	Type       int
	CreateTime time.Time
}

func (v Goods) TableName() string {
	return "goods"
}

func SaveGoods(goods Goods) {
	GetDB().Create(&goods)
}

func UpdateGoods() {
	goods := Goods{}
	GetDB().Where("id = ?", 3).Take(&goods)
	GetDB().Model(&goods).Update("title", "updating")
}

func UpdateGoodsV2() {
	goods := Goods{}
	GetDB().Where("id = ?", 2).Take(&goods)
	//更新非零值的字段  也可以使用map
	GetDB().Model(&goods).Updates(Goods{
		Title: "2",
		Stock: 200,
	})
}

func UpdateGoodsV3() {
	goods := Goods{}
	GetDB().Where("id = ?", 3).Take(&goods)
	GetDB().Model(goods).Select("title").Updates(Goods{
		Title: "3",
		Stock: 300,
	})
}

func UpdateGoodsV4() {
	//g := Goods{}
	//GetDB().Where("id = ?", 4).Take(&g)
	GetDB().Where("id = ?", 4).Omit("title").Updates(Goods{
		Title: "trial",
		Stock: 400,
	})

}

func UpdateGoodsV5() {
	goods := Goods{}
	GetDB().Where("id = ?", 2).Take(&goods)
	//GetDB().Model(&goods).Update("stock", gorm.Expr("stock + 1"))
	GetDB().Model(&goods).Updates(map[string]interface{}{"stock": gorm.Expr("stock + 1")})
}

func UpdateGoodsV6() {
	goods := Goods{}
	GetDB().Where("id = ?", 2).Take(&goods)
	GetDB().Model(&goods).Update("title", GetDB().Model(&User{}).Select("username").Where("id=?", 2))

}

func DeleteGoods() {
	db := GetDB()
	goods := Goods{}
	db.Where("id = ?", 2).Take(&goods)
	db.Delete(&goods)
	//db.Where("id = ?", 2).Delete()

	//根据主键删除
	db.Delete(&Goods{}, 1)
}

func ReadGoods() {
	db := GetDB()
	var goods Goods
	db.Where("title = ? ", "goods").Take(&goods)
	fmt.Println(goods.Id)
}

func ReadGoodsV2() {
	db := GetDB()
	var goods Goods
	db.Where("title = ? ", "goods").First(&goods)
	fmt.Println(goods.Id)
}

func ReadGoodsV3() {
	db := GetDB()
	var goods Goods
	db.Where("title = ? ", "goods").Last(&goods)
	fmt.Println(goods.Id)
}

func ReadGoodsV4() {
	db := GetDB()
	var goods []Goods
	db.Where("title = ? ", "goods").Find(&goods)
	fmt.Println(goods)
}

func ReadGoodsV5() {
	db := GetDB()
	var titles []string
	db.Model(&Goods{}).Pluck("title", &titles)
	fmt.Println(titles)
}

func ReadGoodsV6() {
	db := GetDB()
	var goods Goods
	db.Where("id in (?)", []int{1, 2, 5, 6}).Take(&goods)
	db.Select("id", "title").Find(&goods)
	db.Order("id desc").Find(&goods)
	db.Order("create_time desc").Limit(10).Offset(10).Find(&goods)

	var total int64 = 0
	db.Model(Goods{}).Count(&total)
	fmt.Println(total)

}

type Result struct {
	Type  int
	Total int
}

func ReadGroup() {

	//scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
	//这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定goods表，所以这里只能使用scan查询函数。

	var results []Result
	db := GetDB()
	//等价于: SELECT type, count(*) as  total FROM `goods` GROUP BY type HAVING (total > 0)
	db.Model(Goods{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)
	fmt.Println(results)
}

func ReadRaw() {
	db := GetDB()
	var results []Result

	sql := "SELECT type, count(*) as  total FROM `goods` where create_time > ? GROUP BY type HAVING (total > 0)"
	//因为sql语句使用了一个问号(?)作为绑定参数, 所以需要传递一个绑定参数(Raw第二个参数).
	//Raw函数支持绑定多个参数
	db.Raw(sql, "2022-11-06 00:00:00").Scan(&results)
	fmt.Println(results)
}

// hook
func (Goods) BeforeSave(tx *gorm.DB) error {
	log.Println("saving...")
	return nil
}

func (Goods) BeforeCreate(tx *gorm.DB) error {
	log.Println("creating...")
	return nil
}

func (Goods) AfterCreate(tx *gorm.DB) error {
	log.Println("created...")
	return nil
}

func (Goods) AfterSave(tx *gorm.DB) error {
	log.Println("saved...")
	return nil
}

func (Goods) BeforeUpdate(tx *gorm.DB) error {
	log.Println("updating...")
	return nil
}

func (Goods) AfterUpdate(tx *gorm.DB) error {
	log.Println("updated...")
	return nil
}

func (Goods) BeforeDelete(tx *gorm.DB) error {
	log.Println("deleting...")
	return nil
}

func (Goods) AfterDelete() error {
	log.Println("deleted...")
	return nil
}

func (Goods) AfterFind(tx *gorm.DB) error {
	log.Println("finded...")
	return nil
}
