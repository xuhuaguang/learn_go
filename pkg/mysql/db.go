package mysql

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn_go/config"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	// DB
	DB        *gorm.DB
	initOnce  sync.Once
	ErrDBNull = errors.New("db is null,please connect first")
)

func OpenDB(m *config.MySQLConf, printSql bool) *gorm.DB {
	initOnce.Do(func() {
		config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			m.User,
			m.Password,
			m.Host,
			m.Port,
			m.Name,
		)
		db, err := gorm.Open("mysql", config)
		if err != nil {
			panic(fmt.Sprintf("open db error, %s", err.Error()))
		}
		db.DB().SetConnMaxLifetime(100 * time.Second)
		db.DB().SetMaxIdleConns(10) //最大打开的连接数
		db.DB().SetMaxOpenConns(20)
		db.LogMode(printSql) // 显示详细日志
		DB = db
	})
	return DB
}
