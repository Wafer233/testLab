package gin

import "github.com/gin-gonic/gin"

/*
	1.Tips: RESTful API

	GET，表示读取服务器上的资源 -> read
	POST，表示在服务器上创建资源 -> create
	PUT,表示更新或者替换服务器上的资源 -> update
	DELETE，表示删除服务器上的资源 -> delete
	PATCH，表示更新/修改资源的一部分
*/

/*
	2.req, any, many
*/

func Req() {
	r := gin.Default()
	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, "get")
	})
	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, "post")
	})
	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, "delete")
	})
	r.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(200, "put")
	})

	r.Run(":8080")
}

func ReqAny() {
	r := gin.Default()
	r.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(200, "any")
	})
	r.Run(":8080")
}

func ReqMany() {
	r := gin.Default()
	r.PUT("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "hello,world")
	}).GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "hello,world")
	})
	r.Run(":8080")
}

/*
	3.URIs 3-type
*/

func URLV1() {
	r := gin.Default()

	r.POST("/user/find", func(ctx *gin.Context) {
		ctx.JSON(200, "find nothing")
	})

	r.POST("/user/find/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		ctx.JSON(200, param)
	})

	r.Run(":8080")

}

func URLV2() {
	r := gin.Default()

	r.POST("/user/*path", func(ctx *gin.Context) {
		param := ctx.Param("path")
		ctx.JSON(200, param)
	})

	r.Run(":8080")
}

/*
	4.group
*/

func Group() {
	r := gin.Default()
	ug := r.Group("/user")
	{
		ug.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "user find")
		})
		ug.POST("save", func(ctx *gin.Context) {
			ctx.JSON(200, "user save")
		})
	}
	gg := r.Group("/goods")
	{
		gg.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "goods find")
		})
		gg.POST("save", func(ctx *gin.Context) {
			ctx.JSON(200, "goods save")
		})
	}
	r.Run(":8080")
}
