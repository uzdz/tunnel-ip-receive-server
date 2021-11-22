package entity

type EntryEvent struct {
	DeviceNum  string `json:"deviceNum" bson:"device_num"`   // 设备号
	UId        string `json:"uId" bson:"u_id"`               // 用户ID
	OriginIp   string `json:"originIp" bson:"origin_ip"`     // 源IP
	OriginPort string `json:"originPort" bson:"origin_port"` // 源IP端口
	ProxyIp    string `json:"proxyIp" bson:"proxy_ip"`       // 代理IP
	ProxyPort  string `json:"proxyPort" bson:"proxy_port"`   // 代理端口
	RemoteIp   string `json:"remoteIp" bson:"remote_ip"`     // 远程IP
	RemotePort string `json:"remotePort" bson:"remote_port"` // 远程端口
	EventTime  int32  `json:"eventTime" bson:"event_time"`   // 事件触发事件
	AppId      string `json:"appId" bson:"app_id"`           // 应用Id
}
