package rdbms

import (
	"fmt"
	"ip-receive-server/internal/pkg/core"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	db, err := gorm.Open(core.ApplicationContent.Source.Rdbms.Schema, core.ApplicationContent.Source.Rdbms.Args)
	if err != nil {
		panic(fmt.Sprintf("unable to create rdbms client:  #%v ", err))
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)

	fmt.Printf("Connected to %s!\n", core.ApplicationContent.Source.Rdbms.Schema)

	core.RdbmsContent = NewProduce(db)
}
