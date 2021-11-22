package main

import (
	"flag"
	"fmt"
	http_spec "ip-receive-server/api/http-spec"
	"ip-receive-server/configs"
	"ip-receive-server/internal/pkg/core"
	"ip-receive-server/pkg/sink/kafka"
	"ip-receive-server/pkg/sink/mongodb"
	"ip-receive-server/pkg/source/rdbms"
	"ip-receive-server/plugins/gin_middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "", "配置文件地址")
	flag.Parse()
	if configPath == "" {
		panic(fmt.Sprintf("You must have one config."))
	}

	core.ApplicationContent = configs.GetConfig(configPath)

	if !core.ApplicationContent.Source.Gin.Enable {
		panic(fmt.Sprintf("You must have one source."))
	}

	r := gin.Default()
	r.NoRoute(gin_middleware.Notfound)
	r.Use(gin_middleware.ErrorRecover)
	r.Use(gin_middleware.Cors())
	http_spec.Register(r)
	if core.ApplicationContent.Sink.Kafka.Enable {
		kafka.Init()
	} else if core.ApplicationContent.Sink.Mongodb.Enable {
		mongodb.Init()
	}

	rdbms.Init()

	core.Check()
	defer func() {
		// 关闭资源
		core.Close()
	}()

	port := strconv.Itoa(core.ApplicationContent.Source.Gin.Port)

	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
