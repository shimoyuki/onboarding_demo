package po

import "fmt"

type UserRelationship struct {
	Id                  int
	UserId           string
	FollowUserId     string
	State             int16
}

func (relationship UserRelationship) String() string {
	return fmt.Sprintf("User<%s %s>", relationship.FollowUserId, relationship.State)
}

const (
    DEFAULT = iota
	LIKE
	DISLIKE
	MATCHED
)