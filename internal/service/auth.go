package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"log"
	"net/url"
)

// TODO set redirect URI

const (
	stateSaltLen        = 10
	encryptionKeySize   = 32
	oauthInfoSessionKey = "OAUTH_INFO_SESSION_KEY"
	clientID            = "fancy-rattle-chest"
	clientSecret        = "YOUR_CLIENT_SECRET"
	redirectURI         = "YOUR_REDIRECT_URI"
	authURL             = "https://open-platform-redirect.divar.ir/auth"
	tokenURL            = "https://api.divar.ir/v1/open-platform/oauth/access_token"
)

type OAuth struct {
	AccessToken  string
	RefreshToken string
	Expires      time.Time
	Additional   map[string]interface{}
}

func getOAuth(code string) (*OAuth, error) {
	// OAuth client code to get token (mock implementation)
	token := struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Expires      int64  `json:"expires"`
	}{}

	// Ideally, this would be an actual HTTP request
	tokenJson := `{"access_token": "your_access_token", "refresh_token": "your_refresh_token", "expires": 1700000000}`
	if err := json.Unmarshal([]byte(tokenJson), &token); err != nil {
		return nil, err
	}

	expires := time.Unix(token.Expires, 0)

	return &OAuth{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expires:      expires,
	}, nil
}

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[num.Int64()]
	}
	return string(result)
}

func createRedirectLink(session *sessions.Session, scopes []string, stateData string) (string, error) {
	salt := generateRandomString(stateSaltLen)

	var state string
	if stateData != "" {
		encrypted := encryptStateData(fmt.Sprintf("%s\n%s", stateData, salt))
		state = base64.URLEncoding.EncodeToString(encrypted)
	} else {
		state = salt
	}

	oauthURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&state=%s", authURL, clientID, redirectURI, strings.Join(scopes, " "), state)

	session.Values[oauthInfoSessionKey] = map[string]interface{}{
		"state":     state,
		"scopes":    scopes,
		"oauth_url": oauthURL,
	}

	return oauthURL, nil
}

func encryptStateData(data string) []byte {
	key := []byte("your_32_byte_encryption_key")
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		panic(err)
	}

	// Hash the key to ensure it's the right size for secretbox
	hashKey := sha256.Sum256(key)
	return secretbox.Seal(nonce[:], []byte(data), &nonce, &hashKey)
}

func decryptStateData(state string) (string, error) {
	key := []byte("your_32_byte_encryption_key")
	var nonce [24]byte
	encrypted, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return "", err
	}

	copy(nonce[:], encrypted[:24])
	hashKey := sha256.Sum256(key)
	decrypted, ok := secretbox.Open(nil, encrypted[24:], &nonce, &hashKey)
	if !ok {
		return "", fmt.Errorf("decryption error")
	}

	return string(decrypted[:strings.Index(string(decrypted), "\n")]), nil
}

func createChatSendMessageScope(userId string, postToken string, peerId string) string {
	combined := fmt.Sprintf("%s:%s:%s", userId, postToken, peerId)
	encoded := base64.StdEncoding.EncodeToString([]byte(combined))
	return fmt.Sprintf("CHAT_SEND_MESSAGE_OAUTH__%s", encoded)
}

func getAccessToken(authCode string) string {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	response, err := http.PostForm(tokenURL, data)
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	accessToken := extractAccessToken(body)
	return accessToken
}

func extractAccessToken(responseBody []byte) string {
	var respData struct {
		AccessToken string `json:"access_token"`
		Expires     int    `json:"expires"`
	}

	err := json.Unmarshal(responseBody, &respData)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	return respData.AccessToken
}
