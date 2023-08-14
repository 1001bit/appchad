package blogchad

import (
	"net/http"
)

func PostComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "couldn't parse form", http.StatusBadRequest)
		return
	}
}
