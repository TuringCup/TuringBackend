package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TuringCup/TuringBackend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use DB as database
var Db *gorm.DB

func ConnectDB() {
	DBconfig := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&", DBconfig.UserName, DBconfig.Password, DBconfig.Host, DBconfig.Port, DBconfig.DbName, DBconfig.Charset)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic("connect to db failed,the error is " + err.Error())
	}

	// set connection pool
	sqlDb, err := database.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Minute * 10)

	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDb.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to db successfully")
	Db = database
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := Db
	return db.WithContext(ctx)
}
