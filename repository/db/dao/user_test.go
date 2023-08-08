package dao

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
	"github.com/TuringCup/TuringBackend/repository/db/model"
)

func TestCreateAndDelUser(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	user := model.User{
		Name: "LirdDel", Password: "LirdLirdLird",
		Phone: "13661577632", School: "Njust",
		SchoolID: "9211080N0225", Email: "ruidongli2003@gmail.com",
	}
	res := Db.Create(&user)
	Db.Delete(&user)
	if res.Error != nil {
		t.Error(res.Error)
	}
}
