package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJSON(w, code, map[string]string{"error": msg})
}

func cleanChirp(body string) string {
	words := strings.Split(body, " ")

	for i, word := range words {
		lwrWord := strings.ToLower(word)
		if lwrWord == "kerfuffle" || lwrWord == "sharbert" || lwrWord == "fornax" {
			words[i] = "****"
		}
	}
	body = strings.Join(words, " ")
	return body
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
	} else {
		cleanBody := cleanChirp(params.Body)

		respondWithJSON(w, 200, map[string]string{"cleaned_body": cleanBody})
	}
}
