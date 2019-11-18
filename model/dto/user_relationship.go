package dto

type UserRelationship struct {
	FollowUserId string `json:"user_id"`
	State        string `json:"state"`
	DataType     string `json:"type"`
}
