package core

import (
	"ip-receive-server/configs"
	"ip-receive-server/internal/app/sink"
	"ip-receive-server/internal/app/source"
)

var SinkContent sink.Sink

var RdbmsContent source.Rdbms

var ApplicationContent configs.Content

func Close() {
	SinkContent.Close()
	RdbmsContent.Close()
}

func Check() {

	if SinkContent == nil {
		panic("You must need one sink.")
	}

	if RdbmsContent == nil {
		panic("You must need one rdbms resource.")
	}
}
