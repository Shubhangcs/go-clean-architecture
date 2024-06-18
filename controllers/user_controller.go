package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Shubhangcs/go-clean-architecture/models"
)

type UserHelper struct {
	userRepoInterf models.UserRepository
}

func NewHelperInstance(rep models.UserRepository) *UserHelper {
	return &UserHelper{
		userRepoInterf: rep,
	}
}

func (hlp UserHelper) CreateTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := hlp.userRepoInterf.CreateUserTable()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(models.SimplePayload{Message: "Table Created Successfully"})
}

func (hlp UserHelper) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := io.ReadAll(r.Body)
	var userInfo models.User
	errs := json.Unmarshal(data, &userInfo)
	if errs != nil {
		http.Error(w, "Something Went Wrong", http.StatusBadRequest)
	}
	if err := hlp.userRepoInterf.RegisterUser(userInfo); err != nil {
		http.Error(w, "Register Unsuccessfull", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(models.SimplePayload{Message: "User registeration Successfull"})
}
