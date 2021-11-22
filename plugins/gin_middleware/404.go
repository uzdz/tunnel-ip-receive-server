package gin_middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Notfound(c *gin.Context) {
	log.Println("接口或请求方法找不到：" + c.Request.URL.Path + "、" + c.Request.Method)
	c.JSON(404, gin.H{
		"code": "1004",
		"msg":  "请求地址无效。",
	})
}
