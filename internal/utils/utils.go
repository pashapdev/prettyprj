package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		respondErr := RespondWithError(w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		if respondErr != nil {
			return err
		}
		return fmt.Errorf("failed to marshall payload: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err = w.Write(response); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}
	return nil
}

func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return RespondWithJSON(w,
		code,
		struct {
			Error string `json:"error"`
		}{
			Error: message,
		})
}

func RespondWith400(w http.ResponseWriter, message string) error {
	return RespondWithError(w,
		http.StatusBadRequest,
		message)
}

func RespondWith404(w http.ResponseWriter) error {
	return RespondWithError(w,
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound))
}

func RespondWith500(w http.ResponseWriter) error {
	return RespondWithError(w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError))
}

func SuccessRespondWith200(w http.ResponseWriter, payload interface{}) error {
	return RespondWithJSON(w,
		http.StatusOK,
		payload)
}

func SuccessRepondWith201(w http.ResponseWriter, payload interface{}) error {
	return RespondWithJSON(w,
		http.StatusCreated,
		payload)
}
