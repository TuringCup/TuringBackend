package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SkyAPM/go2sky"
	gormPlugin "github.com/SkyAPM/go2sky-plugins/gorm"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/TuringCup/TuringBackend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use DB as database
var Db *gorm.DB

func ConnectDB() {
	DBconfig := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", DBconfig.UserName, DBconfig.Password, DBconfig.Host, DBconfig.Port, DBconfig.DbName, DBconfig.Charset)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic("connect to db failed,the error is " + err.Error())
	}
	re, err := reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("init tracer error: %v", err)
	}
	defer re.Close()
	// init tracer
	tracer, err := go2sky.NewTracer("TuringDAO", go2sky.WithReporter(re))
	if err != nil {
		log.Fatalf("init tracer error: %v", err)
	}
	database.Use(gormPlugin.New(tracer, gormPlugin.WithPeerAddr(config.Conf.Skywalking.Host+":"+config.Conf.Skywalking.Port), gormPlugin.WithSqlDBType(gormPlugin.MYSQL)))
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

func TestDBClient() *gorm.DB {
	db := Db
	return db
}
