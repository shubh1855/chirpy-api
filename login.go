package main

import (
	"encoding/json"
	"net/http"
	"time"

	"httpserver/internal/auth"

	"github.com/google/uuid"
)

type LoginResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
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

	expiresIn := time.Hour

	if params.ExpiresInSeconds > 0 {
		expiresIn = time.Duration(params.ExpiresInSeconds) * time.Second

		if expiresIn > time.Hour {
			expiresIn = time.Hour
		}
	}

	token, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		expiresIn,
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

	respondWithJSON(w, http.StatusOK, LoginResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	})
}
