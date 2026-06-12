package main

import (
	"encoding/json"
	"net/http"

	"httpserver/internal/auth"
)

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

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
