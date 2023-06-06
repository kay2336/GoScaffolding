package sql

import (
	"awesomeProject/model/table"
	"awesomeProject/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type connMysql struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

var _db *gorm.DB

func InitMysql() {
	// viper 读取yaml配置文件
	conf := utils.GetMysqlByViper()

	// gorm连接mysql
	dsn := conf.DbUser + ":" + conf.DbPassword +
		"@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" +
		conf.DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表不加s
		},
	})
	if err != nil {
		panic(err)
	}

	// 初始化连接池？
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 更新数据库对象
	_db = db

	// 自动建表
	migration()
}

func GetMysqlDB() *gorm.DB {
	return _db
}

func migration() {
	err := _db.AutoMigrate(&table.User{}, &table.Task{})
	if err != nil {
		panic(err)
	}
}
