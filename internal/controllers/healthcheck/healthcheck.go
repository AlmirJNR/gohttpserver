package healthcheck

import (
	"encoding/json"
	"goHttpServer/internal/models/responses"
	"goHttpServer/pkg/utils"
	"net/http"
)

func Handle(w http.ResponseWriter, _ *http.Request) {
	response, _ := json.Marshal(responses.HealthCheck{Status: http.StatusOK})
	_, _ = utils.WriteJson(w, http.StatusOK, response)
}
