package gin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Cookie(r *gin.Engine) {

	r.GET("/set", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("site_cookie", "cookievalue",
			100, "/", "localhost", false, true)
		c.String(200, "cookie has been set")
	})

	r.GET("/read", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("site_cookie")
		if err != nil {
			c.String(200, "not found!")
			return
		}
		c.String(200, data)
	})

	r.GET("/delete", func(c *gin.Context) {
		// 设置cookie  MaxAge设置为-1，表示删除cookie
		c.SetCookie("site_cookie", "cookievalue",
			-1, "/", "localhost", false, true)
		c.String(200, "cookie has been deleted")
	})
}

func Session(r *gin.Engine) {
	// 创建基于cookie的存储引擎，secret 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/session", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		session.Set("hello", "world")
		session.Save()
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
}

func SessionMany(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/sessions", func(c *gin.Context) {
		sessionA := sessions.DefaultMany(c, "a")
		sessionB := sessions.DefaultMany(c, "b")

		sessionA.Set("hello", "world, a-san")
		sessionB.Set("hello", "world, b-san")

		c.JSON(200, gin.H{
			"a": sessionA.Get("hello"),
			"b": sessionB.Get("hello"),
		})
	})
}
