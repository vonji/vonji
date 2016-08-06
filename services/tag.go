package services

import (
	"strings"

	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
)

type TagService struct {
	BaseService
}

func (service TagService) GetAll() []models.Tag {
	if Error != nil {
		return nil
	}

	tags := []models.Tag{}

	if db:= service.GetDB().Find(&tags); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return tags
}

func (service TagService) GetOne(id uint) *models.Tag {
	if Error != nil {
		return nil
	}

	tag := models.Tag{}

	if db := service.GetDB().First(&tag, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return &tag
}

func (service TagService) GetOneWhere(tag *models.Tag) *models.Tag {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(&tag).First(&tag); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return tag
}

func (service TagService) GetAllWhere(tag *models.Tag) []models.Tag {
	if Error != nil {
		return nil
	}

	tags := []models.Tag{}

	if db := service.GetDB().Where(&tag).Find(&tag); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	return tags
}

func (service TagService) GetOrCreate(tags []models.Tag) []models.Tag {
	if  Error != nil {
		return nil
	}

	for i, tag := range tags {
		if tag.ID != 0 {
			continue
		}
		tag.Name = strings.ToLower(tag.Name)
		if t := Tag.GetOneWhere(&tag); t == nil {
			tags[i] = *Tag.Create(&tag)
		} else {
			tags[i] = *t
		}
	}

	return tags
}

func (service TagService) Create(tag *models.Tag) *models.Tag {
	if Error != nil {
		return nil
	}

	tag.Name = strings.ToLower(tag.Name)
	if db := service.GetDB().Create(&tag); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Tag.GetOne(tag.ID)
}

func (service TagService) Update(tag *models.Tag) {
	if Error != nil {
		return
	}

	tag.Name = strings.ToLower(tag.Name)
	if db := service.GetDB().Update(&tag); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service TagService) Delete(id uint) {
	if Error != nil {
		return
	}

	tag := models.Tag{}
	tag.ID = id

	if db := service.GetDB().Delete(&tag); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
}