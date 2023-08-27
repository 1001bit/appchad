package misc

import "net/http"

// get cookie function that makes all stuff easier
func GetCookie(name string, w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return ""
	}
	return cookie.Value
}
