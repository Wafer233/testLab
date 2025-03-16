package dao

import (
	"gorm.io/gorm"
	"time"
)

type UserV2 struct {
	gorm.Model
	Name string
}

// 等效于
type UserV3 struct {
	ID        uint `gorms:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorms:"index"`
	Name      string
}
type Author struct {
	Name  string
	Email string
}

type BlogV1 struct {
	ID      int
	Author  Author `gorms:"embedded"`
	Upvotes int32
}

// 等效于
type BlogV2 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

type BlogV3 struct {
	ID      int
	Author  Author `gorms:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

// 等效于
type BlogV4 struct {
	ID          int64
	AuthorName  string
	AuthorEmail string
	Upvotes     int32
}
