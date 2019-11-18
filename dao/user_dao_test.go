package dao

import (
	"encoding/json"
	"log"
	"onboarding-demo/model/dto"
	"onboarding-demo/model/po"
	"testing"
)

var users = []po.User{
	{
		UserId:   CreateUUID(),
		UserName: "Peter Jones",
		Password: Encrypt("peter_pass"),
		Gender:   1,
	},
	{
		UserId:   CreateUUID(),
		UserName: "John Smith",
		Password: Encrypt("john_pass"),
		Gender:   1,
	},
	{
		UserId:   CreateUUID(),
		UserName: "Susan White",
		Password: Encrypt("susan_pass"),
		Gender:   2,
	},
}

func Test_UserCreate(test *testing.T) {
	_, user := InsertUser(users[0])
	log.Printf("inserted : %s", user)
	_, user = InsertUser(users[0])
	log.Printf("inserted : %s", user)
}

func Test_UserDelete(test *testing.T) {
	//Insert(users[1])
	//log.Printf("deleted : %d", Delete(users[1].UserId))
	log.Printf("deleted : %d", DeleteUser("840555c9-39eb-4bc6-7b3a-6088ce0605c6"))
}

func Test_UserUpdate(test *testing.T) {
	InsertUser(users[2])
	log.Printf("updated : %d", UpdateUser(users[2]))
}

func Test_UserByUUID(test *testing.T) {
	user := UserByUserId("123")
	log.Printf("user : %s", user)
}

func Test_Users(test *testing.T) {
	users := Users()
	for i, user := range users {
		log.Printf("user %d : %s", i, user)
	}
}

func Test_Json(test *testing.T) {
	var userDTO dto.User
	userDTO.UserName = "test"
	body, _ := json.Marshal(&userDTO)
	log.Println(string(body))
	json.Unmarshal([]byte("{\"name\":\"test\"}"), &userDTO)
	log.Println(userDTO)
}
