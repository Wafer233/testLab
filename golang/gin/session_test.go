package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCookie(t *testing.T) {
	r := gin.Default()
	Cookie(r)
	r.Run(":8080")
}

func TestSession(t *testing.T) {
	r := gin.Default()
	Session(r)
	r.Run(":8080")
}

func TestSessionMany(t *testing.T) {
	r := gin.Default()
	SessionMany(r)
	r.Run(":8080")
}
