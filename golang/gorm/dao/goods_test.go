package dao

import (
	"testing"
	"time"
)

func TestSaveGoods(t *testing.T) {
	goods := Goods{
		Title:      "goods",
		Price:      25,
		Stock:      100,
		Type:       0,
		CreateTime: time.Now(),
	}
	SaveGoods(goods)
}

func TestUpdateGoods(t *testing.T) {
	UpdateGoods()
}

func TestUpdateGoodsV2(t *testing.T) {
	UpdateGoodsV2()
}

func TestUpdateGoodsV3(t *testing.T) {
	UpdateGoodsV3()
}

func TestUpdateGoodsV4(t *testing.T) {
	UpdateGoodsV4()
}

func TestUpdateGoodsV5(t *testing.T) {
	UpdateGoodsV5()
}

func TestUpdateGoodsV6(t *testing.T) {
	UpdateGoodsV6()
}

func TestDeleteGoods(t *testing.T) {
	DeleteGoods()
}

func TestReadGoods(t *testing.T) {
	ReadGoods()
}
func TestReadGoodsV2(t *testing.T) {
	ReadGoodsV2()
}
func TestReadGoodsV3(t *testing.T) {
	ReadGoodsV3()
}

func TestReadGoodsV4(t *testing.T) {
	ReadGoodsV4()
}

func TestReadGoodsV5(t *testing.T) {
	ReadGoodsV5()
}

func TestReadGoodsV6(t *testing.T) {
	ReadGoodsV6()
}

func TestReadGroup(t *testing.T) {
	ReadGroup()
}

func TestReadRaw(t *testing.T) {
	ReadRaw()
}
