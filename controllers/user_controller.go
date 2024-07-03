package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shubhangcs/go-clean-architecture/models"
)

type UserController struct {
	userRepositoryInstance models.UserModelInterface
}

func NewUserController(uri models.UserModelInterface) *UserController {
	return &UserController{
		uri,
	}
}

func (uc *UserController) CreateUserTableRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := uc.userRepositoryInstance.CreateUserTable(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Table Created Successfully"})
}

func (uc *UserController) RegisterUserRequest(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	if err := uc.userRepositoryInstance.RegisterUser(&r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "User Registration Successfull"})
}
