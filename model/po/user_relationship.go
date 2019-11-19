package po

import "fmt"

type UserRelationship struct {
	Id           int
	UserId       string
	FollowUserId string
	State        int16
}

func (relationship UserRelationship) String() string {
	return fmt.Sprintf("User<%s %d>", relationship.FollowUserId, relationship.State)
}
