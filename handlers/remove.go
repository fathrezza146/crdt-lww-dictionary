package handlers

import (
	"io"
	"net/http"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	key := r.PostFormValue("key")
	LWW.Remove(key)

	io.WriteString(w, "wew")
}
