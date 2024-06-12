package bean

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var MysqlDb *gorm.DB

func MysqlInit() {
	var dsn = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local"
	dsn = fmt.Sprintf(dsn, Config.Mysql.Username, Config.Mysql.Password, Config.Mysql.Host, Config.Mysql.Port, Config.Mysql.Database, Config.Mysql.ParseTime)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,  // 慢 SQL 阈值
	// 		LogLevel:                  logger.Error, // 日志级别
	// 		IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  true,         // 禁用彩色打印
	// 	},
	// )

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		Log.Panic("连接数据库失败")
	}

	// 实际调整的是 LogLevel
	if Config.Mysql.Debug {
		db = db.Debug()
	}

	MysqlDb = db

	sqlDB, err := db.DB()
	if err != nil {
		Log.Panic("连接数据库失败")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
