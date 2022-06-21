package user

import (
	"github.com/karry-11110/LM_system/app/provider/user"
)

// ConvertUserToDTO 将user转换为UserDTO
func ConvertUserToDTO(user *user.User) *UserDTO {
	if user == nil {
		return nil
	}
	return &UserDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
	}
}
