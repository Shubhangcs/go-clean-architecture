package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shubhangcs/go-clean-architecture/models"
)




type UserHelper struct{
	userRepoInterf models.UserRepository
}

func NewHelperInstance(rep models.UserRepository) *UserHelper{
	return &UserHelper{
		userRepoInterf: rep,
	}
}

func (hlp UserHelper) CreateTable(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	err := hlp.userRepoInterf.CreateUserTable()
	if err != nil{
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(models.SimplePayload{Message: "Someting Went Wrong"})
	}
	json.NewEncoder(w).Encode(models.SimplePayload{Message: "Table Created Successfully"})
}
