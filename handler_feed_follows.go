package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kennethpensopay/rssagg/internal/database"
	"net/http"
	"time"
)

func (a *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	tz, _ := time.LoadLocation("Europe/Copenhagen")
	now := time.Now().In(tz)

	feedFollow, err := a.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (a *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := a.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (a *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")

	if feedFollowId, err := uuid.Parse(feedFollowIDStr); err == nil {
		if err := a.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
			ID:     feedFollowId,
			UserID: user.ID,
		}); err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
			return
		}
	} else {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse Feed Follow ID: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
