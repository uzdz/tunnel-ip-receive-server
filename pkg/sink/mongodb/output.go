package mongodb

import (
	"context"
	"fmt"
	"ip-receive-server/internal/app/sink"
	"ip-receive-server/internal/pkg/core"
	"ip-receive-server/internal/pkg/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	Client   *mongo.Client
	producer *mongo.Database
}

func (s *service) Send(event interface{}) {
	entryEvent := event.(entity.EntryEvent)

	if core.ApplicationContent.Sink.Mongodb.CollectionName == "" {
		return
	}

	collection := s.producer.Collection(core.ApplicationContent.Sink.Mongodb.CollectionName)

	_, err := collection.InsertOne(context.TODO(), entryEvent)
	if err != nil {
		fmt.Sprintf("insert mongodb data failed:  #%v ", err)
	}
}

func (s *service) Close() {
	(*s.Client).Disconnect(context.TODO())
}

func NewProduce(client *mongo.Client, database *mongo.Database) sink.Sink {
	return &service{
		Client:   client,
		producer: database,
	}
}
