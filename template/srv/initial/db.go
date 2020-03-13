package initial

import (
	orm "github.com/eopenio/borm"
	"github.com/eopenio/util/logutil"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true

	err := orm.RegisterDataBase("report", "mysql", "root:@127.0.0.1/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logutil.BgLogger().Info(err.Error())
	}
}
