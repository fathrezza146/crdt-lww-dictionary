package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func Lookup(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	// no input key same as finding empty string key
	lookup := LWW.Lookup(key)
	msg := GeneralResponse{
		Success: true,
		Data:    lookup,
	}

	resp, err := json.Marshal(msg)
	if err != nil {
		io.WriteString(w, "error marshal")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
