package handlers

import (
	"encoding/json"
	"net/http"
)

func internalServerError(w http.ResponseWriter) {
	resp := jsonResponse{
		OK:      false,
		Message: "Internal server error",
	}

	out, _ := json.MarshalIndent(resp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
