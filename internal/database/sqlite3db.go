package database

import (
	"github.com/meloalright/guora/conf"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLITE3DB *gorm.DB
var err error

func init() {

	SQLITE3DB, err = gorm.Open(sqlite.Open(conf.Config().Sql.Addr),&gorm.Config{} )

	_ = SQLITE3DB.Callback().Create().Remove("gorm:update_time_stamp")
	_ = SQLITE3DB.Callback().Update().Remove("gorm:update_time_stamp")

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
}
