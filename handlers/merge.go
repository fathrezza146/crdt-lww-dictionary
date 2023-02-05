package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func Merge(w http.ResponseWriter, r *http.Request) {
	key1 := r.PostFormValue("first_key")
	key2 := r.PostFormValue("second_key")
	if key1 == key2 {
		msg := GeneralResponse{
			Success: false,
			Message: "cannot merge the same keys",
		}

		resp, err := json.Marshal(msg)
		if err != nil {
			io.WriteString(w, "error marshal")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
	LWW.Merge(key1, key2)

	io.WriteString(w, "wew")
}
