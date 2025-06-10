package db

import (
	"go-fiber-starter/internal/model/User"
)

func GetUserById(id string) (model.User, error) {
	var user model.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
