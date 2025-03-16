package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(r *gin.Engine) {

	//string - body
	r.GET("/string", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is a %s", "ms string response")
	})

	type JsonUser struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	//json - body
	ju := JsonUser{
		Id:   1,
		Name: "wafer",
	}
	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ju)
	})

	//xml - body
	type XmlUser struct {
		Id   int64  `xml:"id"`
		Name string `xml:"name"`
	}
	r.GET("/xml", func(ctx *gin.Context) {
		u := XmlUser{
			Id:   1,
			Name: "wafer",
		}
		ctx.XML(http.StatusOK, u)
	})

	//header
	r.GET("/header", func(ctx *gin.Context) {
		ctx.Header("test", "headertest")
	})

	//Redirect
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.google.com")
	})

	r.GET("yaml", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "ms", "age": 19})
	})
}
