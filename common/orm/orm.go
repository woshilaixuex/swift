package orm

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	*gorm.DB
}

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
}

func NewSqlServer(conf *Config) (*DB, error) {
	// 设置默认值
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}
	// 连接数据库
	db, err := gorm.Open(sqlserver.Open(conf.DSN), &gorm.Config{
		Logger: &ormLog{},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetMaxIdleConns(conf.MaxIdleConns)
	sdb.SetMaxOpenConns(conf.MaxOpenConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))
	//err = db.Use(NewCustomePlugin())
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
func MustNewSqlServer(conf *Config) *DB {
	db, err := NewSqlServer(conf)
	if err != nil {
		panic(err)
	}
	return db
}
