package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Temp(r *gin.Engine) {
	r.LoadHTMLFiles("gins/templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// HTML请求
		// 模板的渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "title with template",
		})
	})
}
