package utils

import (
	"net/http"
	"encoding/json"
	"log"
	"errors"

	"go_auth/internal/services"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil{
		log.Println("Failed to parse response data to JSON")
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]interface{}{
		"success": false,
		"error": message,
	})
}

func WriteData(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, map[string]interface{}{
		"success": true,
		"data": data,
	})
}

func ReadJSON(r *http.Request, maxBytes int64, data interface{}) error {
	r.Body = http.MaxBytesReader(nil, r.Body, maxBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(data); err != nil {
		return err
	}

	if dec.More() {
		return errors.New("extra data in request body")
	}
	
	return nil
}

func GetRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return false
	}


	// Limit to 1MB
	if err := ReadJSON(r, (1 << 20), req); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return false
	}

	return true
}

func HandleServiceErrors(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrValidation):
		WriteError(w, http.StatusBadRequest, "emai and password are required")
	case errors.Is(err, services.ErrUserExists):
		WriteError(w, http.StatusConflict, "user with email already exists")
	case errors.Is(err, services.ErrInvalidCredentials):
		WriteError(w, http.StatusUnauthorized, "invalid email or password")
	case errors.Is(err, services.ErrPasswordTooLong):
		WriteError(w, http.StatusBadRequest, "password is too long")
	case errors.Is(err, services.ErrInternal):
		fallthrough
	default:
		WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
