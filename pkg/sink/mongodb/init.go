package mongodb

import (
	"context"
	"fmt"
	"ip-receive-server/internal/pkg/core"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(core.ApplicationContent.Sink.Mongodb.Addrs)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Sprintf("unable to create mongodb client:  #%v ", err))
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Sprintf("unable to create mongodb client:  #%v ", err))
	}
	fmt.Println("Connected to MongoDB!")

	// 指定获取要操作的数据集
	db := client.Database(core.ApplicationContent.Sink.Mongodb.DbName)

	core.SinkContent = NewProduce(client, db)
}
