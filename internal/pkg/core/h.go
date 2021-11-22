package core

import (
	"ip-receive-server/internal/pkg/entity"
	"strings"
)

func Handle(ed entity.EntryEvent) {

	if ed.DeviceNum == "" || ed.EventTime == 0 ||
		ed.RemoteIp == "" || ed.ProxyIp == "" ||
		ed.OriginIp == "" {
		return
	}

	origin := strings.Split(strings.Trim(ed.OriginIp, " "), ":")
	if len(origin) == 2 {
		ed.OriginIp = origin[0]
		ed.OriginPort = origin[1]
	}

	proxy := strings.Split(strings.Trim(ed.ProxyIp, " "), ":")
	if len(proxy) == 2 {
		ed.ProxyIp = proxy[0]
		ed.ProxyPort = proxy[1]
	}

	remote := strings.Split(strings.Trim(ed.RemoteIp, " "), ":")
	if len(remote) == 2 {
		ed.RemoteIp = remote[0]
		ed.RemotePort = remote[1]
	}

	ed.AppId = RdbmsContent.QueryAppIdByVPSNumber(ed.DeviceNum)

	go SinkContent.Send(ed)
}
