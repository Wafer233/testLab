package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetReq(t *testing.T) {
	r := gin.Default()
	GetReq(r)
	r.Run(":8080")
}

func TestGetReqV2(t *testing.T) {
	r := gin.Default()
	GetReqV2(r)
	r.Run(":8080")
}

func TestGetArrayV1(t *testing.T) {
	r := gin.Default()
	GetArrayV1(r)
	r.Run(":8080")
}

func TestGetArrayV2(t *testing.T) {
	r := gin.Default()
	GetArrayV2(r)
	r.Run(":8080")
}

func TestGetMapV1(t *testing.T) {
	r := gin.Default()
	GetMapV1(r)
	r.Run(":8080")
}

func TestPostParam(t *testing.T) {
	r := gin.Default()
	PostParm(r)
	r.Run(":8080")
}

func TestPostParamV2(t *testing.T) {
	r := gin.Default()
	PostParmV2(r)
	r.Run(":8080")
}

func TestPostParamV3(t *testing.T) {
	r := gin.Default()
	PostParmV3(r)
	r.Run(":8080")
}

func TestPostParamV4(t *testing.T) {
	r := gin.Default()
	PostParmV4(r)
	r.Run(":8080")
}
