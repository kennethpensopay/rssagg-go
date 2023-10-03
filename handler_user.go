package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/kennethpensopay/rssagg/internal/database"
	"net/http"
	"time"
)

func (a *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	tz, _ := time.LoadLocation("Europe/Copenhagen")
	now := time.Now().In(tz)

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
