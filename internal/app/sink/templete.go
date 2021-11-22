package sink

type Sink interface {
	Send(event interface{}) // 处理事件
	Close()                 // 关闭资源
}
