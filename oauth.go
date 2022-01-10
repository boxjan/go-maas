package maas

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

// Not a true uuidgen, but at least creates same length random
func generateNonce() string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 48)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

func generateTimestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

type OAuthSigner interface {
	OAuthSign(request *http.Request)
}

type OAuthToken struct {
	ConsumerKey    string
	ConsumerSecret string
	TokenKey       string
	TokenSecret    string
}

// Trick to ensure *plainTextOAuthSigner implements the OAuthSigner interface.
var _ OAuthSigner = (*plainTextOAuthSigner)(nil)

type plainTextOAuthSigner struct {
	token *OAuthToken
	realm string
}

func NewPlainTestOAuthSigner(token *OAuthToken, realm string) (OAuthSigner, error) {
	return &plainTextOAuthSigner{token, realm}, nil
}

// OAuthSign OAuthSignPLAINTEXT signs the provided request using the OAuth PLAINTEXT
// method: http://oauth.net/core/1.0/#anchor22.
func (signer plainTextOAuthSigner) OAuthSign(request *http.Request) {

	signature := signer.token.ConsumerSecret + `&` + signer.token.TokenSecret

	authData := map[string]string{
		"realm":                  signer.realm,
		"oauth_consumer_key":     signer.token.ConsumerKey,
		"oauth_token":            signer.token.TokenKey,
		"oauth_signature_method": "PLAINTEXT",
		"oauth_signature":        signature,
		"oauth_timestamp":        generateTimestamp(),
		"oauth_nonce":            generateNonce(),
		"oauth_version":          "1.0",
	}
	// Build OAuth header.
	var authHeader []string
	for key, value := range authData {
		authHeader = append(authHeader, fmt.Sprintf(`%s="%s"`, key, url.QueryEscape(value)))
	}
	strHeader := "OAuth " + strings.Join(authHeader, ", ")
	request.Header.Add("Authorization", strHeader)
}
