package router

import (
	"github.com/gin-gonic/gin"
	"gos/gorms/api"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
