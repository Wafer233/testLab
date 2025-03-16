package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestResponse(t *testing.T) {
	r := gin.Default()
	Response(r)
	r.Run("8080")
}
