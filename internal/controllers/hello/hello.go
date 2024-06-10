package hello

import (
	"encoding/json"
	"fmt"
	"goHttpServer/internal/models/requests"
	"goHttpServer/internal/models/responses"
	"goHttpServer/pkg/utils"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var request requests.Hello

	decoder := utils.GetDecoder(r.Body)
	err := decoder.Decode(&request)
	errWasHandled := utils.HandleHttpError(w, err)
	if errWasHandled {
		return
	}

	response, err := json.Marshal(responses.Hello{
		Response: fmt.Sprintf("Hello %v", request.Name),
	})
	errWasHandled = utils.HandleHttpError(w, err)
	if errWasHandled {
		return
	}

	_, _ = utils.WriteJson(w, http.StatusOK, response)
}
