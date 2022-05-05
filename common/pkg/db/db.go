package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
	"github.com/Zinces/micro-service/common/pkg/config"
	"fmt"
)

var dbConfig *config.Db
var gormDb *gorm.DB

type BaseModel struct {
	ID        uint64    "gorm:column:id;primaryKey;autoIncrement;not null"
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

func (b *BaseModel) GetStringID() string  {
	return strconv.Itoa(int(b.ID))
}

func (b *BaseModel) CreatedAtDate() string  {
	return b.CreatedAt.Format("2006-01-02 15:04:05")
}

func (b *BaseModel) UpdatedAtDate() string  {
	return b.UpdatedAt.Format("2006-01-02 15:04:05")
}

func connectDB() (*gorm.DB, error)  {
	serviceConfig := config.LoadConfig()
	dbConfig = serviceConfig.Db

	gormDb, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Address, dbConfig.Database, dbConfig.Charset,
	)), &gorm.Config{})

	if err != nil {
		return nil,err
	}

	return gormDb, nil
}

func setupDB()  {
	conn,err := connectDB()
	if err != nil {
		panic(err)
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB")
	conn.Set("gorm:table_options", "Charset=utf8")

	sqlDB,err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("connection to db error %v", err))
	}

	//2.设置最大连接数
	sqlDB.SetMaxOpenConns(dbConfig.MaxConnections)

	//3.设置最大空闲连接数
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdeConnections)

	//4. 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(dbConfig.ConnectionMaxLifeTime * time.Minute)

	//5.设置好连接池，重新赋值
	gormDb = conn
}

func GetDB() *gorm.DB  {
	if gormDb == nil {
		 setupDB()
	}
	return gormDb
}
