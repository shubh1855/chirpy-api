package main

import (
	"httpserver/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Missing refresh  token",
			err,
		)
		return
	}

	err = cfg.db.RevokeRefreshToken(
		r.Context(),
		refreshToken,
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

	w.WriteHeader(http.StatusNoContent)
}
