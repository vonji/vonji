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

	if db:= service.GetDB().Order("id asc").Find(&achievements); db.Error != nil {
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

	if db := service.GetDB().Where(achievement).Order("id asc").Find(&achievements); db.Error != nil {
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

	if db := service.GetDB().Save(&achievement); db.Error != nil {
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

var validation []func(*models.User, *models.Achievement) bool = []func(*models.User, *models.Achievement) bool {
	func(user *models.User, achievement *models.Achievement) bool {
		return false
	},
	func(user *models.User, achievement *models.Achievement) bool {
		requests := Request.GetAllWhere(&models.Request{ Post: models.Post{ UserID: user.ID } })
		return requests != nil && len(requests) >= achievement.CheckData
	},
	func(user *models.User, achievement *models.Achievement) bool { //TODO: refactor when VON-96 is done
		responses := []models.Request{}

		if db := Achievement.GetDB().Where(&models.Response{ Post: models.Post{ UserID: user.ID } }).Not(&models.Response{ Rating: 0 }).Find(&responses); db.Error != nil {
			Error = utils.DatabaseError(db)
			return false
		}
		return len(responses) >= achievement.CheckData
	},
	func(user *models.User, achievement *models.Achievement) bool {
		responses := Response.GetAllWhere(&models.Response{ Post: models.Post{ UserID: user.ID } })
		return responses != nil && len(responses) >= achievement.CheckData
	},
	func(user *models.User, achievement *models.Achievement) bool {
		/*requests := Request.GetAllWhere(&models.Request{ Post: models.Post{ UserID: user.ID } })
		return requests != nil && len(requests) >= achievement.CheckData*/
		return false
	},
	func(user *models.User, achievement *models.Achievement) bool {
		requests := Request.GetAllWhere(&models.Request{ Status: "accepted" })
		tags := Tag.GetAll()
		counts := make([]int, len(tags) + 1)

		for _, request := range requests {
			if Response.GetOneWhere(&models.Response{ Post: models.Post{ UserID: user.ID }, Accepted: true, RequestID: request.ID }) != nil {
				for j := range request.Tags {
					counts[j] += 1
				}
			}
		}
		for _, count := range counts {
			if count >= achievement.CheckData {
				return true
			}
		}

		return false
	},
}

func (service AchievementService) Award() {
	achievements := service.GetAll()
	users := User.GetAll()

	for _, achievement := range achievements {
		if achievement.CheckID != 0 {
			for _, user := range users {
				if !lookatmeimsocoolidontevenneedfunctionalfunctions(user.Achievements, achievement.ID) && validation[achievement.CheckID](&user, &achievement) {
					user.Achievements = append(user.Achievements, achievement)
					user.VActions += achievement.Award
					User.Update(&user)
					Transaction.Create(&models.Transaction{ FromID: 1, ToID: user.ID, Type: "VACTION", Amount: achievement.Award, Reason: "Achievement award", Source: "/achievements" })
				}
				if Error != nil {
					Error = nil
				}
			}
		}
	}
}

func lookatmeimsocoolidontevenneedfunctionalfunctions(achievements []models.Achievement, achievementID uint) bool {
	for _, achievement := range achievements {
		if achievement.ID == achievementID {
			return true
		}
	}
	return false
}