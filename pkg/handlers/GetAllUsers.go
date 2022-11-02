package handlers

import (
	"back/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllUser(w http.ResponseWriter, request *http.Request) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
