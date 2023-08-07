package dao

import (
	"github.com/TuringCup/TuringBackend/repository/db/model"
)

func CreateUser(user model.User) (int32, error) {
	res := Db.Create(&user)
	return user.ID, res.Error
}
