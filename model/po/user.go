package po

import (
	"fmt"
)

type User struct {
	Id       int
	UserId   string
	UserName string
	Password string
	Gender   int16
}

func (user User) String() string {
	return fmt.Sprintf("User<%s %s>", user.UserId, user.UserName)
}
