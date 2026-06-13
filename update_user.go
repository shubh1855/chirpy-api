package main

import (
	"encoding/json"
	"net/http"

	"httpserver/internal/auth"
	"httpserver/internal/database"
)

type updateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
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

	var params updateUserRequest

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't hash password",
			err,
		)
		return
	}

	user, err := cfg.db.UpdateUser(
		r.Context(),
		database.UpdateUserParams{
			ID:             userID,
			Email:          params.Email,
			HashedPassword: hashedPassword,
		},
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't update user",
			err,
		)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
