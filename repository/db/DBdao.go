package db

import (
	"fmt"
	"github.com/TuringCup/TuringBackend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use DB as database
var DB *gorm.DB

func ConnectDB() {
	DBconfig := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&", DBconfig.UserName, DBconfig.Password, DBconfig.Host, DBconfig.Port, DBconfig.DbName, DBconfig.Charset)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to db failed,the error is " + err.Error())
	}
	fmt.Println("connect to db successfully")
	DB = database
}
