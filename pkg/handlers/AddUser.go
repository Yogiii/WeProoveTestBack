package handlers

import (
	"back/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (h handler) AddUser(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User created")
}
