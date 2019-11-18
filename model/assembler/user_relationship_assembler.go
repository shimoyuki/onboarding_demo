package assembler

import (
	"onboarding-demo/model/dto"
	"onboarding-demo/model/po"
)

func ToRelationshipDTO(relationPO po.UserRelationship) (relationDTO dto.UserRelationship) {
	relationDTO.FollowUserId = relationPO.FollowUserId
	switch relationPO.State {
	case po.LIKE:
		relationDTO.State = "like"
	case po.DISLIKE:
		relationDTO.State = "dislike"
	case po.MATCHED:
		relationDTO.State = "matched"
	default:
		relationDTO.State = "default"
	}
	relationDTO.DataType = "relationship"
	return
}

func ToRelationshipDTOs(relationPOs []po.UserRelationship) (relationDTOs []dto.UserRelationship) {
	length := len(relationPOs)
	relationDTOs = make([]dto.UserRelationship, length)
	for i := 0; i < length; i++ {
		relationDTOs[i] = ToRelationshipDTO(relationPOs[i])
	}
	return
}
