package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	list := LWW.List()

	msg := GeneralResponse{
		Success: true,
		Data:    list,
	}

	resp, err := json.Marshal(msg)
	if err != nil {
		io.WriteString(w, "error marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
