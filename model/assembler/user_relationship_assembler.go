package assembler

import (
	"onboarding-demo/model/dto"
	"onboarding-demo/model/enum"
	"onboarding-demo/model/po"
)

const dataTypeRelationship = "relationship"

func ToRelationshipDTO(relationPO po.UserRelationship) (relationDTO dto.UserRelationship) {
	relationDTO.FollowUserId = relationPO.FollowUserId
	relationDTO.State = enum.NameByCode(relationPO.State)
	relationDTO.DataType = dataTypeRelationship
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
