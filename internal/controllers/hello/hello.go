package hello

import (
	"encoding/json"
	"fmt"
	"goHttpServer/internal/models/requests"
	"goHttpServer/internal/models/responses"
	"goHttpServer/pkg/utils"
	"goHttpServer/pkg/utils/guard"
	"goHttpServer/pkg/utils/sanitizer"
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

	errWasHandled = utils.HandleHttpErrors(w,
		guard.AgainstEmptyString(request, "Name"),
		guard.AgainstEmptyString(request, "PhoneNumber"),
	)
	if errWasHandled {
		return
	}

	sanitizedName := sanitizer.SanitizeName(request.Name)
	sanitizedPhoneNumber := sanitizer.SanitizePhoneNumber(request.PhoneNumber)
	response, err := json.Marshal(responses.Hello{
		Response: fmt.Sprintf("Hello %v, can i call you at %v?", sanitizedName, sanitizedPhoneNumber),
	})
	errWasHandled = utils.HandleHttpError(w, err)
	if errWasHandled {
		return
	}

	_, _ = utils.WriteJson(w, http.StatusOK, response)
}
