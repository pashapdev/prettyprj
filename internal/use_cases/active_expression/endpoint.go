package activeexpression

import (
	"net/http"

	"prettyprj/internal/logger"
	"prettyprj/internal/utils"
)

func MakeHandler(s *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := s.Do()
		respondErr := utils.SuccessRespondWith200(w, response)
		if respondErr != nil {
			logger.Error(respondErr)
		}
	}
}
