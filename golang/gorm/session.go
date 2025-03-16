package gorm

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gos/gorms/dao"
	"time"
)

// Session 配置
type Session struct {
	DryRun                   bool //生成 SQL 但不执行
	PrepareStmt              bool //预编译模式
	NewDB                    bool //新db 不带之前的条件
	Initialized              bool //初始化新的db
	SkipHooks                bool //跳过钩子
	SkipDefaultTransaction   bool //禁用默认事务
	DisableNestedTransaction bool //禁用嵌套事务
	AllowGlobalUpdate        bool //允许不带条件的更新
	FullSaveAssociations     bool //允许更新关联数据
	QueryFields              bool //select（字段）
	Context                  context.Context
	Logger                   logger.Interface
	NowFunc                  func() time.Time //允许改变 GORM 获取当前时间的实现
	CreateBatchSize          int
}

func Sessions(db *gorm.DB) {
	var user dao.User
	var users []dao.User
	// 持续会话模式
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	tx.First(&user, 1)
	tx.Find(&users)
	tx.Model(&user).Update("Username", "updating")
}

func SessionsV2(db *gorm.DB) {
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
	tx := db.Session(&gorm.Session{Context: timeoutCtx})

	var user dao.User
	tx.First(&user)                             // 带有 context timeoutCtx 的查询操作
	tx.Model(&user).Update("password", "admin") // 带有 context timeoutCtx 的更新操作
}

func Transactions(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&dao.User{
			Username: "root",
			Password: "root",
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

}

/*
开始事务 				tx := db.Begin()
在事务中执行一些操作	tx.Create(...)
遇到错误时回滚事务		tx.Rollback()
否则，提交事务			tx.Commit()
*/
func TransactionsV2(db *gorm.DB) {
	// 开启事务
	tx := db.Begin()

	err := tx.Model(&dao.Goods{}).Create(&dao.Goods{
		Title:      "new goods",
		Price:      float64(10),
		Stock:      int(100),
		Type:       int(1),
		CreateTime: time.Now(),
	}).Error

	//保存订单失败，则回滚事务
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}
