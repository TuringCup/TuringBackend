package dao

import (
	"fmt"
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

func TestFindUser(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	userdao := TestNewUserDao()
	user, err := userdao.FindUserById(1)
	if err != nil {
		err.Error()
	}
	fmt.Println(*user)
}

func TestUpdateUser(t *testing.T) {
	config.InitConfig("../../..")
	ConnectDB()
	userDao := TestNewUserDao()
	user := model.User{
		Name:     "testName",
		Password: "testPassword",
		Phone:    "testPhone",
		Email:    "test@test",
		School:   "testSchool",
		SchoolID: "testSchoolId",
	}
	success, err := userDao.UpdateUser(1, &user)
	if err != nil {
		err.Error()
	}
	if !success {
		t.Error("update failed")
	}
}
