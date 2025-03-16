package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/*
	1
	query -> ?
	http://localhost:8080/user/save?id=1&name=wafer

	query, defauLtQuery, getQuery
*/

func GetReq(r *gin.Engine) {
	r.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		id, idok := ctx.GetQuery("id")
		address := ctx.DefaultQuery("address", "Tokyo")
		address, aok := ctx.GetQuery("address")
		ctx.JSON(200, gin.H{
			"id":      id,
			"idok":    idok,
			"name":    name,
			"address": address,
			"aok":     aok,
		})
	})
}

/*
form -> uri
json -> json
*/

func GetReqV2(r *gin.Engine) {
	r.GET("/user/save", func(ctx *gin.Context) {
		type User struct {
			Id      int64  `form:"id"`
			Name    string `form:"name"`
			Address string `form:"address" binding:"required"`
		}
		var user User
		err := ctx.BindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
	})
}

// 请求url：http://localhost:8080/user/save?address=Beijing&address=shanghai
func GetArrayV1(r *gin.Engine) {
	r.GET("/user/save", func(ctx *gin.Context) {
		address, ok := ctx.GetQueryArray("address")
		fmt.Println(ok)
		ctx.JSON(200, address)
	})
}

func GetArrayV2(r *gin.Engine) {
	type User struct {
		Id      int64    `form:"id"`
		Name    string   `form:"name"`
		Address []string `form:"address" `
	}
	r.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.BindQuery(&user)
		fmt.Println(err)
		ctx.JSON(200, user)
	})
}

// http://localhost:8080/user/save?addressMap[home]=Beijing&addressMap[company]=shanghai
// map参数 bind并没有支持
func GetMapV1(r *gin.Engine) {
	r.GET("/user/save", func(ctx *gin.Context) {
		addressMap, _ := ctx.GetQueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
}

func PostParm(r *gin.Engine) {
	r.POST("/user/save", func(ctx *gin.Context) {
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")
		address := ctx.PostFormArray("address")
		addressMap := ctx.PostFormMap("addressMap")
		ctx.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"address":    address,
			"addressMap": addressMap,
		})
	})
}
func PostParmV2(r *gin.Engine) {
	r.POST("/user/save", func(ctx *gin.Context) {
		type User struct {
			Id         int64             `form:"id"`
			Name       string            `form:"name"`
			Address    []string          `form:"address"`
			AddressMap map[string]string `form:"addressMap"`
		}
		var user User
		err := ctx.Bind(&user)
		addressMap, _ := ctx.GetPostFormMap("addressMap")
		user.AddressMap = addressMap
		fmt.Println(err)
		ctx.JSON(200, user)
	})
}

func PostParmV3(r *gin.Engine) {
	r.POST("/user/save", func(ctx *gin.Context) {
		type User struct {
			Id         int64             `json:"id"`
			Name       string            `json:"name"`
			Address    []string          `json:"address"`
			AddressMap map[string]string `json:"addressMap"`
		}
		var user User
		err := ctx.BindJSON(&user)
		fmt.Println(err)
		ctx.JSON(200, user)
	})
}

// http://localhost:8080/user/save/111
func PostParmV4(r *gin.Engine) {
	r.POST("/user/save/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})
}
