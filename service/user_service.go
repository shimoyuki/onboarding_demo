package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"onboarding-demo/dao"
	"onboarding-demo/model/assembler"
	"onboarding-demo/model/dto"
	"onboarding-demo/model/enum"
	"onboarding-demo/model/po"
	"onboarding-demo/utils"
)

// GET /users
// query whole user infos
func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(assembler.ToUserDTOs(dao.Users())); err != nil {
		writeErrorResponse(writer, err.Error())
	}
}

func writeErrorResponse(writer http.ResponseWriter, msg string) {
	log.Println(msg)
	writer.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(writer).Encode(msg); err != nil {
		panic(err)
	}
}

// POST /users
// create user
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var userDTO dto.User
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		writeErrorResponse(writer, err.Error())
	}
	defer func() {
		if err := request.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &userDTO); err != nil {
		writeErrorResponse(writer, err.Error())
	}
	if utils.IsStringEmpty(userDTO.UserName) {
		writeErrorResponse(writer, "name can't be blank")
	}
	_, userPO := dao.InsertUser(generateUser(userDTO.UserName))
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(assembler.ToUserDTO(userPO)); err != nil {
		writeErrorResponse(writer, err.Error())
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
	if err := json.NewEncoder(writer).Encode(assembler.ToRelationshipDTOs(dao.RelationshipsByUserId(vars["user_id"]))); err != nil {
		writeErrorResponse(writer, err.Error())
	}
}

// PUT /users/:user_id/relationships/other_user_id
// modify user's relationship
func CreateOrUpdateRelationships(writer http.ResponseWriter, request *http.Request) {
	var relationshipDTO dto.UserRelationship
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		writeErrorResponse(writer, err.Error())
	}
	defer func() {
		if err := request.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &relationshipDTO); err != nil {
		writeErrorResponse(writer, err.Error())
	}
	vars := mux.Vars(request)
	userId, followUserId := vars["user_id"], vars["other_user_id"]
	if utils.IsAnyStringEmpty(userId, followUserId, relationshipDTO.State) {
		writeErrorResponse(writer, "user_id, other_user_id and state can't be blank")
	}
	state := enum.CodeByName(relationshipDTO.State)
	followRelationshipPO := dao.RelationshipsByBothUserIds(followUserId, userId)
	switch state {
	case enum.Like.Code():
		if followRelationshipPO.State == enum.Like.Code() {
			state = enum.Matched.Code()
			followRelationshipPO.State = state
		}
	case enum.Dislike.Code():
		if followRelationshipPO.State == enum.Matched.Code() {
			followRelationshipPO.State = enum.Like.Code()
		}
	default:
		writeErrorResponse(writer, "only state within {\"like\",\"dislike\"} is allowed")
	}
	dao.UpdateRelationship(followRelationshipPO)
	_, relationshipPO := dao.InsertRelationship(generateRelationship(userId, followUserId, state))
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(assembler.ToRelationshipDTO(relationshipPO)); err != nil {
		writeErrorResponse(writer, err.Error())
	}
}

func generateRelationship(userId string, followUserId string, state int16) (relationship po.UserRelationship) {
	relationship = *new(po.UserRelationship)
	relationship.UserId = userId
	relationship.FollowUserId = followUserId
	relationship.State = state
	return
}
