package http_spec

import (
	"ip-receive-server/pkg/source/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// 初始页面
	r.GET("/", http.Welcome)

	// 埋点收集API
	r.POST("/e", http.SourceWithHttp)
}
