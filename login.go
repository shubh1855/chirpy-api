package main

import (
	"encoding/json"
	"net/http"
	"time"

	"httpserver/internal/auth"
	"httpserver/internal/database"

	"github.com/google/uuid"
)

type LoginResponse struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	IsChirpyRed  bool      `json:"is_chirpy_red"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {

	var params loginRequest

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	user, err := cfg.db.GetUserByEmail(
		r.Context(),
		params.Email,
	)

	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Incorrect email or password",
			err,
		)
		return
	}

	valid, err := auth.CheckPasswordHash(
		params.Password,
		user.HashedPassword,
	)

	if err != nil || !valid {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Incorrect email or password",
			err,
		)
		return
	}

	token, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		time.Hour,
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't create token",
			err,
		)
		return
	}

	refreshToken := auth.MakeRefreshToken()

	_, err = cfg.db.CreateRefreshToken(
		r.Context(),
		database.CreateRefreshTokenParams{
			Token:     refreshToken,
			UserID:    user.ID,
			ExpiresAt: time.Now().UTC().Add(60 * 24 * time.Hour),
		},
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't create refresh token",
			err,
		)
	}

	respondWithJSON(w, http.StatusOK, LoginResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		IsChirpyRed:  user.IsChirpyRed,
		Token:        token,
		RefreshToken: refreshToken,
	})
}
