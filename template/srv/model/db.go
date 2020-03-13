package model

import (
	orm "github.com/eopenio/gorm"
	_ "github.com/eopenio/gorm/dialects/mysql"
	"github.com/eopenio/util/terror"
	"time"
)

const (
	DefaultMaxIdleConns    = 10
	DefaultMaxOpenConns    = 20
	DefaultConnMaxLifetime = 1 * time.Hour
)

// 数据库实例
var (
	db *orm.DB
)

func Init() {
	db, err := orm.Open("mysql", "root:@127.0.0.1/test?charset=utf8mb4&parseTime=True&loc=Local")
	terror.MustNil(err)

	// SetMaxIdleCons 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(DefaultMaxIdleConns)

	// SetMaxOpenCons 设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(DefaultMaxOpenConns)

	// SetConnMaxLifetiment 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(DefaultConnMaxLifetime)

	db.LogMode(true)

	// 禁用默认表名的复数形式
	db.SingularTable(true)
}
