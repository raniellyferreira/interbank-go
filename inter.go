package inter

import (
	"context"
	"os"

	"github.com/raniellyferreira/interbank-go/auth"
	"github.com/raniellyferreira/interbank-go/backend"
	"github.com/raniellyferreira/interbank-go/banking"
	"github.com/raniellyferreira/interbank-go/cobranca"
	"github.com/raniellyferreira/interbank-go/pix"
)

// Client is a client for the Inter service
type Client struct {
	// Backend backend
	backend *backend.BackendImplement

	// Conta corrente (opcional)
	accountNumber string

	// Serviços
	Pix      *pix.Service
	Banking  *banking.Service
	Cobranca *cobranca.Service
}

// NewClientWithCredentials creates a new client with the given credentials
func NewClientWithCredentials(creds *auth.Credentials) *Client {
	// Create the backend
	backend := backend.NewBackendWithCredentials(creds)
	backend.SetURL("https://cdpj.partners.bancointer.com.br")

	return &Client{
		backend: backend,

		Pix:      pix.NewService(backend),
		Banking:  banking.NewService(backend),
		Cobranca: cobranca.NewService(backend),
	}
}

// NewClient creates a new client with default credentials loaded from the environment variables (see NewDefaultCredentials)
func NewClient() (*Client, error) {
	creds, err := auth.NewDefaultCredentials()
	if err != nil {
		return nil, err
	}

	// Create the client
	client := NewClientWithCredentials(creds)

	// Check if we should use the sandbox
	if os.Getenv("INTERBANK_USE_SANDBOX") == "true" {
		client.UseSandBox()
	}

	return client, nil
}

// Token returns the current token or fetches a new one if it's expired (thread-safe)
func (c *Client) Token(ctx context.Context) (*auth.Token, error) {
	return c.backend.Token(ctx)
}

// UseSandBox sets the base URL to the sandbox environment (set URL to https://cdpj-sandbox.partners.uatinter.co)
func (c *Client) UseSandBox() *Client {
	c.SetURL("https://cdpj-sandbox.partners.uatinter.co")
	return c
}

// SetURL sets the base URL for the client
func (c *Client) SetURL(url string) *Client {
	c.backend.SetURL(url)
	return c
}

// SetAccountNumber sets the account number for the client
func (c *Client) SetAccountNumber(accountNumber string) *Client {
	c.accountNumber = accountNumber
	c.backend.SetHeader("x-conta-corrente", accountNumber)
	return c
}
