package rdbms

import (
	"ip-receive-server/internal/app/source"

	"github.com/jinzhu/gorm"
)

type service struct {
	db *gorm.DB
}

func (s *service) QueryAppIdByVPSNumber(number string) string {
	var appId string
	rows, _ := s.db.Raw("select pp.app_id as appId from proxy_pool as pp inner join proxy_vps as pv on pp.id = pv.pool_id where pv.id = ?", number).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&appId)
	}
	return appId
}

func (s *service) Close() {
	(*s.db).Close()
}

func NewProduce(db *gorm.DB) source.Rdbms {
	return &service{
		db: db,
	}
}
