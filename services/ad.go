package services

import (
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type AdService struct {
	BaseService
}

func (service AdService) GetAll() []models.Ad {
	if Error != nil {
		return nil
	}

	ads := []models.Ad{}

	if db:= service.GetDB().Find(&ads); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return ads
}

func (service AdService) GetOne(id uint) *models.Ad {
	if Error != nil {
		return nil
	}

	ad := models.Ad{}

	if db := service.GetDB().First(&ad, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return &ad
}

func (service AdService) GetOneWhere(ad *models.Ad) *models.Ad {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(&ad).First(&ad); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return ad
}

func (service AdService) GetAllWhere(ad *models.Ad) []models.Ad {
	if Error != nil {
		return nil
	}

	ads := []models.Ad{}

	if db := service.GetDB().Where(ad).Find(&ads); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return ads
}

func (service AdService) Create(ad *models.Ad) *models.Ad {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&ad); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Ad.GetOne(ad.ID)
}

func (service AdService) Update(ad *models.Ad) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&ad); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service AdService) Delete(id uint) {
	if Error != nil {
		return
	}

	ad := models.Ad{}
	ad.ID = id

	if db := service.GetDB().Delete(&ad); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
}