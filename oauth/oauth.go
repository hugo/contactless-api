package oauth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	googleOAuthCodeEndpoint  = "https://accounts.google.com/o/oauth2/v2/auth"
	googleOAuthTokenEndpoint = "https://www.googleapis.com/oauth2/v4/token"
	redirectURI              = "http://localhost:3000/auth/callback"
	responseType             = "code"
	scopes                   = "http://www.google.com/m8/feeds/contacts/"
	grantType                = "authorization_code"
)

// Token is an oAuth token from Google's oAuth2 server
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int32  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// GetGoogleOAuthURI builds the URL we send the client to for them
// to generate us an access token
func GetGoogleOAuthURI(clientID string) (string, error) {
	uri, err := url.Parse(googleOAuthCodeEndpoint)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Add("response_type", responseType)
	q.Add("client_id", clientID)
	q.Add("redirect_uri", redirectURI)
	q.Add("scope", scopes)
	uri.RawQuery = q.Encode()

	return uri.String(), nil
}

// GetGoogleOAuthAccessToken exchanges an auth code for an access token
func GetGoogleOAuthAccessToken(clientID, clientSecret, code string) (*Token, error) {
	data := url.Values{}
	data.Add("code", code)
	data.Add("client_id", clientID)
	data.Add("client_secret", clientSecret)
	data.Add("redirect_uri", redirectURI)
	data.Add("grant_type", grantType)

	req, err := http.NewRequest(
		"POST",
		googleOAuthTokenEndpoint,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error performing request %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != 200 {
		var errorResponse struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&errorResponse)
		return nil, errors.New(errorResponse.ErrorDescription)
	}

	defer res.Body.Close()
	var token Token
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
