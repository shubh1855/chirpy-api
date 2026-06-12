package main

import (
	"encoding/json"
	"httpserver/internal/auth"
	"httpserver/internal/database"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type createChirpParams struct {
	Body string `json:"body"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func cleanProfanity(text string) string {
	badWords := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	words := strings.Split(text, " ")

	for i, word := range words {
		if badWords[strings.ToLower(word)] {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	var params createChirpParams

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Missing or invalid token",
			err,
		)
		return
	}

	userID, err := auth.ValidateJWT(
		tokenString,
		cfg.jwtSecret,
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Invalid token",
			err,
		)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Invalid request",
			err,
		)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Chirp is too long",
			nil,
		)
		return
	}

	cleanedBody := cleanProfanity(params.Body)

	chirp, err := cfg.db.CreateChirp(
		r.Context(),
		database.CreateChirpParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Body:      cleanedBody,
			UserID:    userID,
		},
	)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't create chirp",
			err,
		)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
