package main

// Spotify Oauth flow
// https://developer.spotify.com/documentation/web-api/tutorials/code-flow

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

const authUrl = "https://accounts.spotify.com/authorize"
const callbackUrl = "https://localhost:3420/callback"
const scope = "user-read-currently-playing"
const showDialog = true // Force approval of auth each run
const clientIdFileName = "clientid"

type AuthData struct {
	authUrl      string
	clintId      string
	responseType string
	redirectUrl  string
	state        string
	scope        string
	showDialog   string
}

func NewAuthData() (AuthData, error) {

	state, err := GenerateNonce()
	if err != nil {
		return AuthData{}, fmt.Errorf("failed to generate state token: %s", err)

	}

	clientId, clientErr := ReadClientId()
	if clientErr != nil {
		return AuthData{}, fmt.Errorf("failed to read clientid file: %s", clientErr)
	}

	data := &AuthData{authUrl, clientId, "code", callbackUrl, state, scope, strconv.FormatBool(showDialog)}
	return *data, nil
}

func GenerateNonce() (string, error) {
	nonceBytes := make([]byte, 32)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", fmt.Errorf("failed to read from crypto/rand")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes), nil
}

func ReadClientId() (string, error) {
	file, err := os.Open(clientIdFileName)
	if err != nil {
		return "", fmt.Errorf("unable to open clientid file, %s", err)
	}
	defer file.Close()

	fileBytes := make([]byte, 32)

	if _, err := file.Read(fileBytes); err != nil {
		return "", fmt.Errorf("unable to read clientif file, %s", err)
	}

	return string(fileBytes), nil
}

func main() {
	authData, err := NewAuthData()
	if err != nil {
		fmt.Printf("Failed generating auth data: %s\n", err)
	}
	fmt.Printf("Auth data: %s\n", authData)

	// response, err := http.Get("https://google.com")
	// if err != nil {
	// 	fmt.Printf("Error making request: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Response status code: %s\n", response.Status)

	// resultBody, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Printf("Could not read the response body!\n")
	// 	os.Exit(1)
	// }
	// fmt.Printf("Response Body: %s\n", resultBody)
}
