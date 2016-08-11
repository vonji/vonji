package services

import (
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type TransactionService struct {
	BaseService
}

func (service TransactionService) GetAll() []models.Transaction {
	if Error != nil {
		return nil
	}

	transactions := []models.Transaction{}

	if db := service.GetDB().Find(&transactions); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, transaction := range transactions {
		transactions[i].From = *User.GetOne(transaction.FromID)
		transactions[i].To = *User.GetOne(transaction.ToID)
	}

	return transactions
}

func (service TransactionService) GetAllWhere(transaction *models.Transaction) []models.Transaction {
	if Error != nil {
		return nil
	}

	transactions := []models.Transaction{}

	if db := service.GetDB().Where(&transaction).Find(&transactions); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, transaction := range transactions {
		transactions[i].From = *User.GetOne(transaction.FromID)
		transactions[i].To = *User.GetOne(transaction.ToID)
	}

	return transactions
}

func (service TransactionService) GetOne(id uint) *models.Transaction {
	if Error != nil {
		return nil
	}

	transaction := models.Transaction{}

	if db := service.GetDB().First(&transaction, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	transaction.From = *User.GetOne(transaction.FromID)
	transaction.To = *User.GetOne(transaction.ToID)

	return &transaction
}

func (service TransactionService) GetOneWhere(transaction *models.Transaction) *models.Transaction {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(&transaction).First(&transaction); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	transaction.From = *User.GetOne(transaction.FromID)
	transaction.To = *User.GetOne(transaction.ToID)

	return transaction
}

func (service TransactionService) Create(transaction *models.Transaction) *models.Transaction{
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&transaction); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Transaction.GetOne(transaction.ID)
}


func (service TransactionService) Update(transaction *models.Transaction) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&transaction); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service TransactionService) Delete(id uint) {
	if Error != nil {
		return
	}

	transaction := models.Transaction{}
	transaction.ID = id

	if db := service.GetDB().Delete(&transaction); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}
