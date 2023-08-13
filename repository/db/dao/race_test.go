package dao

import (
	"github.com/TuringCup/TuringBackend/config"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"testing"
)

func TestCreateAndDelRace(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	race := model.Race{
		Name: "test",
	}
	res := Db.Create(&race)
	if res.Error != nil {
		t.Error(res.Error)
	}
	res = Db.Delete(&race)
	if res.Error != nil {
		t.Error(res.Error)
	}
}
