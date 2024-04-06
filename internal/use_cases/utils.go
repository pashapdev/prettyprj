package usecases

import (
	"encoding/json"
	"net/http"

	"prettyprj/internal/logger"
	"prettyprj/internal/utils"
)

func DecodeBody(w http.ResponseWriter, r *http.Request, request interface{}) error {
	const notValidBodyMessage = "failed to decode request"
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(notValidBodyMessage, err)
		respondErr := utils.RespondWith400(w, notValidBodyMessage)
		if respondErr != nil {
			logger.Error("failed to write response", err)
		}
		return err
	}
	return nil
}
