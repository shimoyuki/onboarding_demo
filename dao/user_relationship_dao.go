package dao

import (
	"log"
	"onboarding-demo/model/po"
)

func InsertRelationship(relationship po.UserRelationship) (bool, po.UserRelationship) {
	isSuccess, err := Db.Model(&relationship).
		Where("id = ?", relationship.Id).
		OnConflict("ON CONSTRAINT uniq_user_follow DO UPDATE").
		Set("state = EXCLUDED.state").
		Returning("id").
		SelectOrInsert()
	if err != nil {
		log.Println(err)
		return false, relationship
	}
	return isSuccess, relationship
}

func DeleteRelationship(id string) int {
	var relationship po.UserRelationship
	result, err := Db.Model(&relationship).
		Where("id = ?", id).
		Delete()
	if err != nil {
		log.Println(err)
		return 0
	}
	return result.RowsAffected()
}

func UpdateRelationship(relationship po.UserRelationship) int {
	result, err := Db.Model(&relationship).
		Where("user_id = ? and follow_user_id = ?", relationship.UserId, relationship.FollowUserId).
		Update()
	if err != nil {
		log.Println(err)
		return 0
	}
	return result.RowsAffected()
}

func RelationshipsByUserId(userId string) (relationships []po.UserRelationship) {
	err := Db.Model(&relationships).
		Where("user_id = ?", userId).
		Select()
	if err != nil {
		log.Println(err)
	}
	return
}

func RelationshipsByBothUserIds(userId string, followUserId string) (relationships po.UserRelationship) {
	err := Db.Model(&relationships).
		Where("user_id = ? and follow_user_id = ?", userId, followUserId).
		Select()
	if err != nil {
		log.Println(err)
	}
	return
}
