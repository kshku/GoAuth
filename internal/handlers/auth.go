package handlers

import (
	"net/http"

	"go_auth/internal/services"
	"go_auth/internal/utils"
)

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"name"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if !utils.GetRequest(w, r, &req) {
		return
	}

	data := services.RegisterUserData{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	response, err := services.RegisterNewUser(r.Context(), &data)
	if err != nil {
		utils.HandleServiceErrors(w, err)
		return
	}

	utils.WriteData(w, http.StatusCreated, map[string]interface{}{
		"user": map[string]string{
			"name": response.Name,
			"email": response.Email,
		},
		"token": response.Token,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if !utils.GetRequest(w, r, &req) {
		return
	}

	data := services.LoginUserData{
		Email: req.Email,
		Password: req.Password,
	}

	response, err := services.LoginUser(r.Context(), &data)
	if err != nil {
		utils.HandleServiceErrors(w, err)
		return
	}

	utils.WriteData(w, http.StatusOK, map[string]interface{}{
		"user": map[string]string{
			"name": response.Name,
			"email": response.Email,
		},
		"token": response.Token,
	})
}
