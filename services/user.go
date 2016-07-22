package services

import (
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type UserService struct {
	BaseService
}

func (service UserService) GetAll() []models.User {
	if Error != nil {
		return nil
	}

	users := []models.User{}

	if err := service.GetDB().Find(&users); err.Error != nil {
		Error = utils.DatabaseError(err)
		return nil
	}

	for i, user := range users {
		if err := service.GetDB().Model(&user).Association("tags").Find(&users[i].Tags); err.Error != nil {
			Error = utils.AssociationError(err)
			return nil
		}
	}

	return users
}

func (service UserService) GetOne(id uint) models.User {
	if Error != nil {
		return models.User{}
	}

	user := models.User{}

	if err := service.GetDB().First(&user, id); err.Error != nil {
		Error = utils.DatabaseError(err)
		return models.User{}
	}

	if err := service.GetDB().Model(&user).Association("tags").Find(&user.Tags); err.Error != nil {
		Error = utils.AssociationError(err)
		return models.User{}
	}

	return user
}

func (service UserService) GetOneByEmail(email string) models.User {
	if Error != nil {
		return models.User{}
	}

	user := models.User{ Email: email }

	if err := service.GetDB().Where(&user).First(&user); err.Error != nil {
		Error = utils.DatabaseError(err)
		return models.User{}
	}

	if err := service.GetDB().Model(&user).Association("tags").Find(&user.Tags); err.Error != nil {
		Error = utils.AssociationError(err)
		return models.User{}
	}

	return user
}

func (service UserService) Create(user models.User) {
	if Error != nil {
		return
	}

	if err := service.GetDB().Create(&user); err.Error != nil {
		Error = utils.DatabaseError(err)
		return
	}
}


func (service UserService) Update(user models.User) {
	if Error != nil {
		return
	}

	if err := service.GetDB().Save(&user); err != nil {
		Error = utils.DatabaseError(err)
	}
}

func (service UserService) Delete(id uint) {
	if Error != nil {
		return
	}

	user := models.User{}
	user.ID = id

	if err := service.GetDB().Delete(&user); err != nil {
		Error = utils.DatabaseError(err)
	}
}