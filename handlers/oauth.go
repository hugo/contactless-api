package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugo/contactless/oauth"
)

// OAuthURI generates a URI where an auth code can be acquired
func OAuthURI(clientID string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		googleOAuthURI, err := oauth.GetGoogleOAuthURI(clientID)
		if err != nil {
			log.Fatalln(
				"Failed to generate Google oAuth URI: %s",
				err.Error(),
			)
			http.Error(
				rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}

		res := struct {
			Ok  bool   `json:"ok"`
			URI string `json:"uri"`
		}{
			true,
			googleOAuthURI,
		}
		json.NewEncoder(rw).Encode(res)
	}
}

// HandleGoogleOAuthCallback receives an auth code from Google
// and turns it into an access token
func HandleGoogleOAuthCallback(clientID, clientSecret string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var data struct {
			Code string `json:"code"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("Error parsing code from request %s\n", err.Error())
			http.Error(
				rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		resp, err := oauth.GetGoogleOAuthAccessToken(
			clientID,
			clientSecret,
			data.Code,
		)
		if err != nil {
			log.Printf("Getting OAuth token failed: %s", err.Error())
			http.Error(
				rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		json.NewEncoder(rw).Encode(resp)
	}
}
