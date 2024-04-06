package healthz

import (
	"net/http"

	"prettyprj/internal/entities"
	"prettyprj/internal/logger"
	"prettyprj/internal/utils"
)

type response struct {
	Name         string `json:"name"`
	BuildVersion string `json:"build_version"`
	BuildTime    string `json:"build_time"`
	GitTag       string `json:"git_tag"`
	GitHash      string `json:"git_hash"`
}

func MakeHandler(info *entities.AppInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		response := &response{
			Name:         info.Name,
			BuildVersion: info.BuildVersion,
			BuildTime:    info.BuildTime,
			GitTag:       info.GitTag,
			GitHash:      info.GitHash,
		}
		err := utils.SuccessRespondWith200(w, struct{}{})
		if err != nil {
			logger.Error("failed to decode response", response, err)
		}
	}
}
