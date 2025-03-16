package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func Middle() {
	r := gin.Default()
	r.Use(gin.Logger())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	r.Use(gin.Recovery())
	r.GET("/middleware", func(ctx *gin.Context) {
		panic(errors.New("middleware use success"))
	})
	r.Run("8080")

}
