package utility

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/entity"
	"github.com/subosito/gotenv"
)

func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		// load .env
		_ = gotenv.Load()
		requiredUser := os.Getenv("BASIC_USERNAME")
		requiredPassword := os.Getenv("BASIC_PASSWORD")

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			SendErrorResponse(w, entity.InvalidTokenError)
		}
	}
}
