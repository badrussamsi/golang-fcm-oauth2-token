package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

type tokenProvider struct {
	tokenSource oauth2.TokenSource
}

// newTokenProvider function to get token for fcm-send
func newTokenProvider(credentialsLocation string) (*tokenProvider, error) {
	jsonKey, err := ioutil.ReadFile(credentialsLocation)
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
	tp, err := newTokenProvider("./private_key/myserversite-6678e-firebase-adminsdk-9bfn4-25a9e89012.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	token, err := tp.token()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(token)

}
