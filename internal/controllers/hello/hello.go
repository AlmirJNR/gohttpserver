package hello

import (
	"encoding/json"
	"fmt"
	"goHttpServer/internal/models/requests"
	"goHttpServer/internal/models/responses"
	"goHttpServer/internal/pkg"
	"goHttpServer/internal/pkg/guard"
	"goHttpServer/internal/pkg/sanitizer"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var request requests.Hello

	decoder := pkg.GetDecoder(r.Body)
	err := decoder.Decode(&request)
	errWasHandled := pkg.HandleHttpError(w, err)
	if errWasHandled {
		return
	}

	errWasHandled = pkg.HandleHttpErrors(w,
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
	errWasHandled = pkg.HandleHttpError(w, err)
	if errWasHandled {
		return
	}

	_, _ = pkg.WriteJson(w, http.StatusOK, response)
}
