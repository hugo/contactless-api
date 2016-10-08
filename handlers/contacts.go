package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hugo/contactless/contacts"
)

var noTokenError = errors.New("No access_token provided")

func getToken(r *http.Request) (string, error) {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return "", noTokenError
	}
	parts := strings.Split(bearer, "Bearer ")
	if len(parts) != 2 {
		return "", noTokenError
	}
	token := parts[1]
	if token == "" {
		return "", noTokenError
	}
	return token, nil
}

// HandleContacts gets contacts from the Google API and
// returns them as JSON
func HandleContacts(rw http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		http.Error(
			rw,
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		)
		return
	}

	req, _ := http.NewRequest(
		"GET",
		googleEndpoint,
		nil,
	)
	req.Header.Add("GData-Version", "3.0")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/atom+xml")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch contacts: %s\n", err.Error())
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if res.StatusCode != 200 {
		http.Error(
			rw,
			http.StatusText(res.StatusCode),
			res.StatusCode,
		)
		return
	}

	defer res.Body.Close()
	type Entries []contacts.Contact
	var feed struct {
		XMLName xml.Name `xml:"feed"`
		Entries `xml:"entry"`
	}
	err = xml.NewDecoder(res.Body).Decode(&feed)
	if err != nil {
		log.Printf("Failed to create contact: %s\n", err.Error())
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	var people struct {
		Contacts []contacts.JSONContact `json:"contacts"`
	}

	for _, entry := range feed.Entries {
		person := contacts.ToJSON(entry)

		people.Contacts = append(
			people.Contacts,
			person,
		)
	}

	rw.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(rw)
	enc.SetIndent("", "  ")
	enc.Encode(people)
}

// ContactAdd adds a contact
func ContactAdd(rw http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		http.Error(
			rw,
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		)
		return
	}

	if r.Method != "POST" {
		http.Error(
			rw,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
		return
	}

	var jsonContact contacts.JSONContact
	json.NewDecoder(r.Body).Decode(&jsonContact)

	contact := contacts.FromJSON(jsonContact)
	googleContact := contacts.Googlify(contact)

	data, err := xml.Marshal(googleContact)
	if err != nil {
		log.Printf("Error generating XML: %s\n", err.Error())
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	req, _ := http.NewRequest(
		"POST",
		googleEndpoint,
		strings.NewReader(string(data)),
	)
	req.Header.Add("GData-Version", "3.0")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/atom+xml")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to add contact: %s\n", err.Error())
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if res.StatusCode != 201 {
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	// TODO: return the newly added contact

	var newContact contacts.Contact
	xml.NewDecoder(res.Body).Decode(&newContact)
	newJSONContact := contacts.ToJSON(newContact)

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(newJSONContact)
}

// ContactDelete removes a contact
func ContactDelete(rw http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		http.Error(
			rw,
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		)
		return
	}

	if r.Method != "DELETE" {
		http.Error(
			rw,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
		return
	}

	var data struct {
		URI  string `json:"uri"`
		Etag string `json:"etag"`
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	// Google is dumb as a box of frogs and gives us a http URI
	uri, _ := url.Parse(data.URI)
	if uri.Scheme == "http" {
		uri.Scheme = "https"
	}

	req, _ := http.NewRequest(
		"DELETE",
		uri.String(),
		nil,
	)
	req.Header.Add("GData-Version", "3.0")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("If-Match", data.Etag)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to delete contact: %s\n", err.Error())
		http.Error(
			rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if res.StatusCode != 200 {
		log.Println(res.StatusCode)
		http.Error(
			rw,
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		)
		return
	}
}
