package services

import (
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
)

type NotificationService struct {
	BaseService
}

func (service NotificationService) GetAll() []models.Notification {
	if Error != nil {
		return nil
	}

	notifications := []models.Notification{}

	if db := service.GetDB().Find(&notifications); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, notification := range notifications {
		notifications[i].User = *User.GetOne(notification.UserID)
	}

	return notifications
}

func (service NotificationService) GetOne(id uint) *models.Notification {
	if Error != nil {
		return nil
	}

	notification := models.Notification{}

	if db:= service.GetDB().First(&notification, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	notification.User = *User.GetOne(notification.UserID)

	return &notification
}

func (service NotificationService) GetOneWhere(notification *models.Notification) *models.Notification {
	if Error != nil {
		return nil
	}

	if db:= service.GetDB().Where(&notification).First(&notification); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	notification.User = *User.GetOne(notification.UserID)

	return notification
}

func (service NotificationService) GetAllWhere(notification *models.Notification) []models.Notification {
	if Error != nil {
		return nil
	}

	notifications := []models.Notification{}

	if db := service.GetDB().Where(&notification).Find(&notifications); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, notification := range notifications {
		notifications[i].User = *User.GetOne(notification.UserID)
	}

	return notifications
}

func (service NotificationService) Create(notification *models.Notification) *models.Notification {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&notification); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Notification.GetOne(notification.ID)
}

func (service NotificationService) Update(notification *models.Notification) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&notification); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service NotificationService) Delete(id uint) {
	if Error != nil {
		return
	}

	notification := models.Notification{}
	notification.ID = id

	if db := service.GetDB().Delete(&notification); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
}

