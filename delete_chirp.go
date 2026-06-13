package main

import (
	"database/sql"
	"httpserver/internal/auth"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
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

	chirpID, err := uuid.Parse(
		r.PathValue("chirpID"),
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Invalid chirp ID",
			err,
		)
		return
	}

	chirp, err := cfg.db.GetChirp(
		r.Context(),
		chirpID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't get chirp",
			err,
		)
		return
	}

	if chirp.UserID != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = cfg.db.DeleteChirp(
		r.Context(),
		chirpID,
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't delete chirp",
			err,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
