package response

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSONWithPayload(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
