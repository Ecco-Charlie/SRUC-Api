package middleware

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/pkg"
)

func AlrredyAuth(next config.PageHandle) config.PageHandle {
	return func(w http.ResponseWriter, r *http.Request) (string, any) {
		jwt, err := r.Cookie("sid")
		if err != nil {
			return next(w, r)
		}
		err = pkg.ValidateJwt(&jwt.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return "", nil
		}
		return next(w, r)
	}
}

func AuthSessionKeyMiddleware(next config.PageHandle) config.PageHandle {
	return func(w http.ResponseWriter, r *http.Request) (string, any) {
		jwt, err := r.Cookie("sid")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return "", nil
		}
		err = pkg.ValidateJwt(&jwt.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return "", nil
		}
		return next(w, r)
	}
}
