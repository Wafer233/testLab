package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestTemp(t *testing.T) {
	r := gin.Default()
	Temp(r)
	r.Run(":8080")
}
