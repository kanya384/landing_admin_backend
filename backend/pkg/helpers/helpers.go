package helpers

import (
	"net/http"
	"time"
)

func SetValueToCookies(w http.ResponseWriter, name string, value string, ttl time.Duration, secured bool) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(ttl),
		Secure:  secured,
	})
	return
}
