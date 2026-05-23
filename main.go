package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

type tokenProvider struct {
	tokenSource oauth2.TokenSource
}

// newTokenProvider function to get token for fcm-send
func newTokenProvider(credentialsLocation string) (*tokenProvider, error) {
	jsonKey, err := os.ReadFile(credentialsLocation)
	if err != nil {
		return nil, errors.New("fcm: failed to read credentials file at: " + credentialsLocation)
	}
	cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
	if err != nil {
		return nil, errors.New("fcm: failed to get JWT config for the firebase.messaging scope")
	}
	ts := cfg.TokenSource(context.Background())
	return &tokenProvider{
		tokenSource: ts,
	}, nil
}

func (src *tokenProvider) token() (string, error) {
	token, err := src.tokenSource.Token()
	if err != nil {
		return "", errors.New("fcm: failed to generate Bearer token")
	}
	return token.AccessToken, nil
}

func main() {
	_ = godotenv.Load()

	// credentialsPath points to the Firebase service account JSON file used to sign OAuth2 tokens for FCM.
	// Set FIREBASE_CREDENTIAL_PATH in .env (or system environment), for example: ./private_key/<service-account>.json
	credentialsPath := os.Getenv("FIREBASE_CREDENTIAL_PATH")
	if credentialsPath == "" {
		log.Fatal("missing FIREBASE_CREDENTIAL_PATH (set it in environment or .env)")
	}

	tp, err := newTokenProvider(credentialsPath)
	if err != nil {
		log.Fatal("file reading error: ", err)
	}

	token, err := tp.token()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(token)

}
