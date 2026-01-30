package handlers

import (
	"net/http"

	"go_auth/internal/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.WriteData(w, http.StatusOK, map[string]string {
		"status": "OK",
	})
}
