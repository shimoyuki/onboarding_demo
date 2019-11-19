package enum

type UserRelationshipEnum struct {
	code int16
	name string
}

var Default = UserRelationshipEnum{code: 0, name: "default"}
var Like = UserRelationshipEnum{code: 1, name: "like"}
var Dislike = UserRelationshipEnum{code: 2, name: "dislike"}
var Matched = UserRelationshipEnum{code: 3, name: "matched"}

func (enum UserRelationshipEnum) Code() int16 {
	switch enum {
	case Default:
		return Default.code
	case Like:
		return Like.code
	case Dislike:
		return Dislike.code
	case Matched:
		return Matched.code
	default:
		panic("Could not map enum")
	}
}

func CodeByName(name string) int16 {
	switch name {
	case Default.name:
		return Default.code
	case Like.name:
		return Like.code
	case Dislike.name:
		return Dislike.code
	case Matched.name:
		return Matched.code
	default:
		panic("Could not map enum")
	}
}

func (enum UserRelationshipEnum) Name() string {
	switch enum {
	case Default:
		return Default.name
	case Like:
		return Like.name
	case Dislike:
		return Dislike.name
	case Matched:
		return Matched.name
	default:
		panic("Could not map enum")
	}
}

func NameByCode(code int16) string {
	switch code {
	case Default.code:
		return Default.name
	case Like.code:
		return Like.name
	case Dislike.code:
		return Dislike.name
	case Matched.code:
		return Matched.name
	default:
		panic("Could not map enum")
	}
}

func Values() []UserRelationshipEnum {
	return []UserRelationshipEnum{
		Default,
		Like,
		Dislike,
		Matched,
	}
}
