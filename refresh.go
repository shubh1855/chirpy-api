package main

import (
	"net/http"
	"time"

	"httpserver/internal/auth"
)

type refreshResponse struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Missing refresh token",
			err,
		)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(
		r.Context(),
		refreshToken,
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Invalid refresh token",
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
			"Couldn't create access token",
			err,
		)
		return
	}

	respondWithJSON(
		w,
		http.StatusOK,
		refreshResponse{
			Token: token,
		},
	)
}
