package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func WriteERROR(w http.ResponseWriter, statusCode int, err error) {
	msgs := struct {
		Error string `json:"error"`
	}{
		Error: http.StatusText(statusCode),
	}

	if err != nil {
		msgs.Error = err.Error()
	}

	WriteJSON(w, statusCode, msgs)
}
