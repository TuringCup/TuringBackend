package dao

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
	"github.com/TuringCup/TuringBackend/repository/db/model"
)

func TestCreateUser(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	_, err := CreateUser(model.User{
		Name: "Lird", Password: "LirdLirdLird",
		Phone: "13661577631", School: "Njust",
		SchoolID: "9211080N0225", Email: "ruidongli2002@gmail.com",
	})
	if err != nil {
		t.Error(err)
	}
}
