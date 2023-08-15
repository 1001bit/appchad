package blogchad

import (
	"net/http"
)

func CommentPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "couldn't parse form", http.StatusBadRequest)
		return
	}
}
