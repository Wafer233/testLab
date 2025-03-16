package gin

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.POST("hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})
	r.Run(":8080")
}
