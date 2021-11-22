package http

import (
	"errors"
	"ip-receive-server/internal/pkg/core"
	"ip-receive-server/internal/pkg/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func contentTypeToJson(c *gin.Context, ed *entity.EntryEvent) error {
	if err := c.ShouldBindJSON(ed); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func SourceWithHttp(c *gin.Context) {
	var ed entity.EntryEvent

	if e := contentTypeToJson(c, &ed); e != nil {
		expError(c, e.Error())
		return
	}

	go core.Handle(ed)

	c.JSON(http.StatusOK, gin.H{
		"code": "ok",
	})
}
