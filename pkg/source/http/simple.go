package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome")
}

func expError(c *gin.Context, msg string) {
	var errorMsg = msg + " invalid."
	fmt.Println(errorMsg)
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
			"code": errorMsg,
		})
}
