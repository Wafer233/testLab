package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.POST("save", SaveUser)
	r.POST("get", GetUser)
	r.POST("update", UpdateUser)
	r.POST("delete", DeleteUser)
}
