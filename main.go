package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hugo/contactless/handlers"
	"github.com/hugo/contactless/middleware"
	"github.com/joho/godotenv"
)

const (
	contactsBaseURI    = "https://www.google.com/m8/feeds/contacts"
	atomSuppliesDomain = "atomsupplies.com"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		log.Fatalln("You must set the environment variable GOOGLE_CLIENT_ID")
	}

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if clientID == "" {
		log.Fatalln("You must set the environment variable GOOGLE_CLIENT_SECRET")
	}

	// log.Printf("Client ID: %s\nClient secret: %s\n", clientID, clientSecret)

	corsConfig := middleware.CORSConfig{
		AllowedCredentials: true,
		AllowedHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins:     []string{"*"},
	}

	http.HandleFunc("/contacts",
		middleware.Chain(handlers.HandleContacts,
			middleware.CORSMiddleware(corsConfig),
		),
	)

	http.HandleFunc("/contacts/delete",
		middleware.Chain(handlers.ContactDelete,
			middleware.CORSMiddleware(corsConfig),
		),
	)

	http.HandleFunc("/contacts/add",
		middleware.Chain(handlers.ContactAdd,
			middleware.CORSMiddleware(corsConfig),
		),
	)

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/auth",
		middleware.Chain(handlers.OAuthURI(clientID),
			middleware.CORSMiddleware(corsConfig),
		),
	)

	http.HandleFunc("/auth/token",
		middleware.Chain(handlers.HandleGoogleOAuthCallback(clientID, clientSecret),
			middleware.CORSMiddleware(corsConfig),
		),
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("You must set the environment variable PORT")
	}

	var listenOn string
	if appEnv == "production" {
		listenOn = "0.0.0.0:" + port
	} else {
		listenOn = "127.0.0.1:" + port
	}
	log.Println("Listening on ", listenOn)
	http.ListenAndServe(listenOn, nil)
}
