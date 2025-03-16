package gorm

import (
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	r := gin.Default()
	db := InitDB()
	Create(db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
