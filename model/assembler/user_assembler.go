package assembler

import (
	"onboarding-demo/model/dto"
	"onboarding-demo/model/po"
)

const dataTypeUser = "user"

func ToUserDTO(userPO po.User) (userDTO dto.User) {
	userDTO.UserId = userPO.UserId
	userDTO.UserName = userPO.UserName
	userDTO.DataType = dataTypeUser
	return
}

func ToUserDTOs(userPOs []po.User) (userDTOs []dto.User) {
	length := len(userPOs)
	userDTOs = make([]dto.User, length)
	for i := 0; i < length; i++ {
		userDTOs[i] = ToUserDTO(userPOs[i])
	}
	return
}
