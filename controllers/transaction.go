package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/vonji/vonji-api/services"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
	"github.com/gorilla/mux"
)

type TransactionController struct {
	APIBaseController
}

func (ctrl TransactionController) GetAll() (interface{}, *utils.HttpError) {
	return services.Transaction.GetAll(), nil
}

func (ctrl TransactionController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Transaction.GetOne(id), nil
}

func (ctrl TransactionController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	transaction := models.Transaction{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &transaction); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Transaction.GetOneWhere(&transaction), nil
}

func (ctrl TransactionController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	transaction := models.Transaction{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &transaction); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Transaction.GetAllWhere(&transaction), nil
}

func (ctrl TransactionController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	transaction := &models.Transaction{}

	if err := json.NewDecoder(r.Body).Decode(transaction); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	if transaction = services.Transaction.Create(transaction); transaction == nil {
		return nil, nil
	}

	if transaction.Type == "VACTION" {
		transaction.To.VAction += transaction.Amount
		services.User.Update(&transaction.To)
	} else if transaction.Type == "VCOIN" {
		transaction.From.VCoins -= transaction.Amount//TODO check if enough coins
		transaction.To.VCoins += transaction.Amount
		services.User.Update(&transaction.From)
		services.User.Update(&transaction.To)

	}
	hammer := "vCoins"
	if transaction.Type == "VACTION" {
		hammer = "vActions"
	}
	go (func() {
		services.Notification.Create(&models.Notification {
			UserID: transaction.FromID,
			Title: fmt.Sprintf("Vous avez été débité de %d %s", transaction.Amount, hammer),
			Message: transaction.Reason,
		})
	})()
	go (func() {
		services.Notification.Create(&models.Notification {
			UserID: transaction.ToID,
			Title: fmt.Sprintf("Vous avez recu %d %s", transaction.Amount, hammer),
			Message: transaction.Reason,
		})
	})()

	return services.Transaction.GetOne(transaction.ID), nil
}

func (ctrl TransactionController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	transaction := models.Transaction{}

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Transaction.Update(&transaction)

	return nil
}

func (ctrl TransactionController) Delete(id uint) *utils.HttpError {
	services.Transaction.Delete(id)
	return nil
}


