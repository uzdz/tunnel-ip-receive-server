package gin_middleware

import (
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ErrorRecover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			// 封装通用json返回
			c.JSON(500, gin.H{
				"code": "9999",
				"msg":  "服务器开小差了～",
			})
		}
	}()

	//加载完 defer recover，继续后续接口调用
	c.Next()
}
