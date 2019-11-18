package dao

import (
	"log"
	"onboarding-demo/model/po"
)

func InsertUser(user po.User) (bool, po.User) {
	created, err := Db.Model(&user).
		Where("user_id = ?", user.UserId).
		OnConflict("DO NOTHING").
		Returning("id").
		SelectOrInsert()
	if err != nil {
		log.Println(err)
		return false, user
	}
	return created, user
}

func DeleteUser(userId string) int {
	var user po.User
	result, err := Db.Model(&user).
		Where("user_id = ?", userId).
		Delete()
	if err != nil {
		log.Println(err)
		return 0
	}
	return result.RowsAffected()
}

func UpdateUser(user po.User) int {
	result, err := Db.Model(&user).
		OnConflict("DO NOTHING").
		Where("user_id = ?", user.UserId).
		Update()
	if err != nil {
		log.Println(err)
		return 0
	}
	return result.RowsAffected()
}

func Users() (users []po.User) {
	err := Db.Model(&users).Select()
	if err != nil {
		log.Println(err)
	}
	return
}

func UserByUserId(userId string) (user po.User) {
	err := Db.Model(&user).
		Where("user_id = ?", userId).
		Select()
	if err != nil {
		log.Println(err)
	}
	return
}
