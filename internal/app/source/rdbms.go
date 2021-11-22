package source

type Rdbms interface {
	QueryAppIdByVPSNumber(number string) string // 根据VPS设备号查询应用Id
	Close()                                     // 关闭资源
}
