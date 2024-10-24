package backend

import (
	"context"
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/raniellyferreira/interbank-go/auth"
	"github.com/raniellyferreira/interbank-go/erros"
)

const (
	oauthEndpoint = "oauth/v2"
)

// type Backend interface {
// 	SetURL(url string) Backend
// 	SetHeader(header, value string) Backend
// 	Req() *resty.Request
// 	Token(ctx context.Context) (*auth.Token, error)
// }

type BackendImplement struct {
	client *resty.Client
	creds  *auth.Credentials

	token *auth.Token
	tmu   sync.Mutex
}

// NewBackendWithCredentials creates a new backend with the given credentials
func NewBackendWithCredentials(creds *auth.Credentials) *BackendImplement {
	// Create the client
	client := resty.New()

	// Set TLS
	if tls := creds.GetTLS(); tls != nil {
		client.SetTLSClientConfig(tls)
	}

	return &BackendImplement{
		client: client,
		creds:  creds,
	}
}

// SetURL sets the base URL for the backend
func (c *BackendImplement) SetURL(url string) *BackendImplement {
	c.client.SetBaseURL(url)
	return c
}

// SetHeader sets a header for the backend
func (c *BackendImplement) SetHeader(header, value string) *BackendImplement {
	c.client.SetHeader(header, value)
	return c
}

// Req get request
func (c *BackendImplement) Req() *resty.Request {
	return c.client.R()
}

// Token returns the current token or requests a new one
func (ts *BackendImplement) Token(ctx context.Context) (*auth.Token, error) {
	ts.tmu.Lock()
	defer ts.tmu.Unlock()

	if ts.token == nil || !ts.token.Valid() {
		token, err := ts.requestNewToken(ctx)
		if err != nil {
			return nil, err
		}

		ts.token = token

		// Set the expiration time
		ts.token.SetExpiresAt(time.Now().Add(time.Duration(ts.token.ExpiresIn-180) * time.Second))
	}

	return ts.token, nil
}

// requestNewToken requests a new token
func (ts *BackendImplement) requestNewToken(ctx context.Context) (*auth.Token, error) {
	// Send the request
	resp, err := ts.Req().
		SetContext(ctx).
		SetResult(&auth.Token{}).
		SetFormData(ts.creds.BuildAuthFormData()).
		Post(path.Join(oauthEndpoint, "token"))
	if err != nil {
		return nil, erros.NewFromError(err)
	}

	// Check for errors
	if resp.IsError() {
		return nil, erros.NewErrorWithStatus(resp.StatusCode(), resp.String())
	}

	token, ok := resp.Result().(*auth.Token)
	if !ok {
		return nil, erros.NewErrorWithStatus(resp.StatusCode(), fmt.Sprintf("invalid response type: %T", resp.Result()))
	}

	return token, nil
}
