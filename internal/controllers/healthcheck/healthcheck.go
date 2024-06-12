package healthcheck

import (
	"encoding/json"
	"goHttpServer/internal/models/responses"
	"goHttpServer/internal/pkg"
	"net/http"
)

func Handle(w http.ResponseWriter, _ *http.Request) {
	response, _ := json.Marshal(responses.HealthCheck{Status: http.StatusOK})
	_, _ = pkg.WriteJson(w, http.StatusOK, response)
}
