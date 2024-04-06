package calcexpression

import (
	"errors"
	"net/http"

	appErrors "prettyprj/internal/errors"
	"prettyprj/internal/logger"
	usecases "prettyprj/internal/use_cases"
	"prettyprj/internal/utils"
)

func MakeHandler(s *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(request)
		err := usecases.DecodeBody(w, r, request)
		if err != nil {
			logger.Error("failed to decode body", err)
			if err = utils.RespondWith400(w, "failed to decode body"); err != nil {
				logger.Error(err)
			}
			return
		}
		defer r.Body.Close()

		err = s.Do(request)
		if err != nil {
			if errors.Is(err, appErrors.ErrUnsupportedOperation) {
				if err = utils.RespondWith400(w, err.Error()); err != nil {
					logger.Error(err)
				}
				return
			} else {
				if err = utils.RespondWith500(w); err != nil {
					logger.Error(err)
				}
				return
			}
		}

		respondErr := utils.SuccessRepondWith201(w, struct{}{})
		if respondErr != nil {
			logger.Error(respondErr)
		}
	}
}
