package services

import (
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type AchievementService struct {
	BaseService
}

func (service AchievementService) GetAll() []models.Achievement {
	if Error != nil {
		return nil
	}

	achievements := []models.Achievement{}

	if db:= service.GetDB().Find(&achievements); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return achievements
}

func (service AchievementService) GetOne(id uint) *models.Achievement {
	if Error != nil {
		return nil
	}

	achievement := models.Achievement{}

	if db := service.GetDB().First(&achievement, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return &achievement
}

func (service AchievementService) GetOneWhere(achievement *models.Achievement) *models.Achievement {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(&achievement).First(&achievement); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return achievement
}

func (service AchievementService) GetAllWhere(achievement *models.Achievement) []models.Achievement {
	if Error != nil {
		return nil
	}

	achievements := []models.Achievement{}

	if db := service.GetDB().Where(achievement).Find(&achievements); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return achievements
}

func (service AchievementService) Create(achievement *models.Achievement) *models.Achievement {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&achievement); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Achievement.GetOne(achievement.ID)
}

func (service AchievementService) Update(achievement *models.Achievement) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Update(&achievement); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service AchievementService) Delete(id uint) {
	if Error != nil {
		return
	}

	achievement := models.Achievement{}
	achievement.ID = id

	if db := service.GetDB().Delete(&achievement); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
}