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

//func TestFindRace(t *testing.T) {
//	config.InitConfig("../../..")
//	ConnectDB()
//	race := model.Race{
//		Name: "test",
//	}
//	racedao := TestNewRaceDao()
//	racedao.CreateRace(&race)
//	raceFind, err := racedao.FindRaceById(1)
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//	fmt.Println(raceFind)
//	//Db.Delete(&race)
//}
