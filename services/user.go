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

	if db := service.GetDB().Find(&users); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, user := range users {
		if db := service.GetDB().Model(&user).Association("tags").Find(&users[i].Tags); db.Error != nil {
			Error = utils.AssociationError(db)
			return nil
		}
	}

	return users
}

func (service UserService) GetOne(id uint) *models.User {
	if Error != nil {
		return nil
	}

	user := models.User{}

	if db := service.GetDB().First(&user, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	if db := service.GetDB().Model(&user).Association("tags").Find(&user.Tags); db.Error != nil {
		Error = utils.AssociationError(db)
		return nil
	}

	return &user
}

func (service UserService) GetOneByEmail(email string) *models.User {
	if Error != nil {
		return nil
	}

	user := models.User{ Email: email }

	if db := service.GetDB().Where(&user).First(&user); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	if db := service.GetDB().Model(&user).Association("tags").Find(&user.Tags); db.Error != nil {
		Error = utils.AssociationError(db)
		return nil
	}

	return &user
}

func (service UserService) Create(user *models.User) *models.User{
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&user); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return User.GetOne(user.ID)
}


func (service UserService) Update(user *models.User) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&user); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service UserService) Delete(id uint) {
	if Error != nil {
		return
	}

	user := models.User{}
	user.ID = id

	if db := service.GetDB().Delete(&user); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}