package dao

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
)

func TestConnection(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	sqlDB, err := Db.DB()
	if err != nil {
		t.Error(err)
	}
	if err = sqlDB.Ping(); err != nil {
		t.Error(err)
	}
}
