package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"onboarding-demo/dao"
	"onboarding-demo/model/assembler"
	"onboarding-demo/model/dto"
	"onboarding-demo/model/po"
)

// GET /users
// query whole user infos
func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(assembler.ToUserDTOs(dao.Users()))
	if err != nil {
		panic(err)
	}
}

// POST /users
// create user
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var userDTO dto.User
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &userDTO); err != nil || userDTO.UserName == "" {
		writer.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}
	_, userPO := dao.InsertUser(generateUser(userDTO.UserName))
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(assembler.ToUserDTO(userPO)); err != nil {
		panic(err)
	}
}

func generateUser(userName string) (user po.User) {
	user = *new(po.User)
	user.UserId = dao.CreateUUID()
	user.UserName = userName
	user.Gender = 0
	user.Password = dao.Encrypt(userName)
	return
}

// GET /users/:user_id/relationships
// query relationships of a user
func GetRelationshipsByUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(assembler.ToRelationshipDTOs(dao.RelationshipsByUserId(vars["user_id"])))
	if err != nil {
		panic(err)
	}
}

// PUT /users/:user_id/relationships/other_user_id
// modify user's relationship
func CreateOrUpdateRelationships(writer http.ResponseWriter, request *http.Request) {
	var relationshipDTO dto.UserRelationship
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &relationshipDTO); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}
	vars := mux.Vars(request)
	userId, followUserId := vars["user_id"], vars["other_user_id"]
	var state int16
	switch relationshipDTO.State {
	case "like":
		followRelationship := dao.RelationshipsByBothUserIds(followUserId, userId)
		if followRelationship.State == po.LIKE {
			state = po.MATCHED
			followRelationship.State = state
			dao.UpdateRelationship(followRelationship)
		} else {
			state = po.LIKE
		}
	case "dislike":
		state = po.DISLIKE
	default:
		state = po.DEFAULT
	}
	_, relationshipPO := dao.InsertRelationship(generateRelationship(userId, followUserId, state))
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(assembler.ToRelationshipDTO(relationshipPO)); err != nil {
		panic(err)
	}
}

func generateRelationship(userId string, followUserId string, state int16) (relationship po.UserRelationship) {
	relationship = *new(po.UserRelationship)
	relationship.UserId = userId
	relationship.FollowUserId = followUserId
	relationship.State = state
	return
}
