package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type AddForm struct {
	Element string "json:element"
}

func Add(w http.ResponseWriter, r *http.Request) {
	key := r.PostFormValue("key")
	value := r.PostFormValue("value")

	LWW.Add(key, value)

	msg := GeneralResponse{
		Success: true,
		Message: "data is added",
	}

	resp, err := json.Marshal(msg)
	if err != nil {
		io.WriteString(w, "error marshal")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
