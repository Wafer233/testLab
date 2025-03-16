package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestReq(t *testing.T) {
	r := gin.Default()
	Req()
	r.Run(":8080")
}

func TestReqAny(t *testing.T) {
	r := gin.Default()
	ReqAny()
	r.Run(":8080")
}

func TestReqMany(t *testing.T) {
	r := gin.Default()
	ReqMany()
	r.Run(":8080")
}

func TestURLV1(t *testing.T) {
	r := gin.Default()
	URLV1()
	r.Run(":8080")
}

func TestURLV2(t *testing.T) {
	r := gin.Default()
	URLV2()
	r.Run(":8080")
}

func TestGroup(t *testing.T) {
	r := gin.Default()
	Group()
	r.Run(":8080")
}
