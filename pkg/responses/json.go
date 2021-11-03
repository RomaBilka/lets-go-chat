package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	msgs := struct {
		Error string `json:"error"`
	}{
		Error: http.StatusText(http.StatusBadRequest),
	}

	if statusCode != 0 {
		msgs.Error = http.StatusText(statusCode)
		if err != nil {
			msgs.Error = err.Error()
		}

		JSON(w, statusCode, msgs)

		return
	}

	JSON(w, http.StatusBadRequest, nil)
}
