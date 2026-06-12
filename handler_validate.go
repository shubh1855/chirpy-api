package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type parameters struct {
	Body string `json:"body"`
}

type cleanedResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	var params parameters

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
		return
	}

	cleaned := cleanProfanity(params.Body)

	respondWithJSON(w, http.StatusOK, cleanedResponse{
		CleanedBody: cleaned,
	})
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
