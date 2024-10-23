package auth

import (
	"crypto/tls"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Credentials struct {
	clientID     string  // Client Id obtido no detalhe da tela de aplicações no IB
	clientSecret string  // Client Secret obtido no detalhe da tela de aplicações no IB
	grantType    string  // GrantType que utilizamos, o default é (client_credentials)
	scopes       []Scope // Scopes

	cert *tls.Certificate // TLS
}

func NewCredentials(clientID, clientSecret string, scopes ...Scope) *Credentials {
	return &Credentials{
		clientID:     clientID,
		clientSecret: clientSecret,
		scopes:       scopes,
		grantType:    "client_credentials",
	}
}

// BuildAuthFormData returns the form data for the credentials
func (c *Credentials) BuildAuthFormData() map[string]string {
	data := map[string]string{
		"client_id":     c.clientID,
		"client_secret": c.clientSecret,
		"grant_type":    c.grantType,
	}

	if c.HasScopes() {
		data["scope"] = c.GetScopesString()
	}

	return data
}

/*
Set the following environment variables to configure the default credentials:

INTERBANK_CLIENT_ID
and
INTERBANK_CLIENT_SECRET

INTERBANK_SCOPES - must be a comma-separated list of scopes

INTERBANK_TLS_PATH - (required for mutual TLS)
must be a path to a file containing the certificate and key (tls.key and tls.crt)
*/
func NewDefaultCredentials() (*Credentials, error) {
	// Get environment variables
	clientID := os.Getenv("INTERBANK_CLIENT_ID")
	clientSecret := os.Getenv("INTERBANK_CLIENT_SECRET")

	// Check if the client ID and secret are set
	if clientID == "" || clientSecret == "" {
		return nil, errors.New("INTERBANK_CLIENT_ID and INTERBANK_CLIENT_SECRET must be set")
	}

	// Create credentials
	creds := NewCredentials(clientID, clientSecret)

	// Set scopes
	creds.SetScopesFromString(os.Getenv("INTERBANK_SCOPES"))

	// Load TLS files
	tlsPath := os.Getenv("INTERBANK_TLS_PATH")
	if tlsPath != "" {
		err := creds.LoadCertAndKeyFromPath(filepath.Join(tlsPath, "tls.crt"), filepath.Join(tlsPath, "tls.key"))
		if err != nil {
			return nil, err
		}
	}

	return creds, nil
}

// GetTLS returns the TLS configuration
func (c *Credentials) GetTLS() *tls.Config {
	if c.cert == nil {
		return nil
	}

	return &tls.Config{
		Certificates: []tls.Certificate{*c.cert},
	}
}

// SetTLS sets the TLS certificate and key
func (c *Credentials) SetTLS(cert, key []byte) *Credentials {
	c.cert = &tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  key,
	}
	return c
}

// LoadCertAndKeyFromPath loads the TLS certificate and key from the given paths
func (c *Credentials) LoadCertAndKeyFromPath(certPath, keyPath string) error {
	// Load the certificate and key
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return err
	}

	c.cert = &cert
	return nil
}

// HasScopes returns true if the credentials have scopes
func (c *Credentials) HasScopes() bool {
	return len(c.scopes) > 0
}

// GetScopes returns the scopes
func (c *Credentials) GetScopes() []Scope {
	return c.scopes
}

// GetScopesString returns the scopes as a space-separated string
func (c *Credentials) GetScopesString() string {
	scopes := make([]string, len(c.scopes))
	for i, scope := range c.scopes {
		scopes[i] = string(scope)
	}
	return strings.Join(scopes, " ")
}

/*
SetScopesFromString sets the scopes from a comma-separated string

Escopos cadastrados na tela de aplicações.

Escopos disponíveis:

extrato.read - Consulta de Extrato e Saldo
boleto-cobranca.read - Consulta de boletos e exportação para PDF
boleto-cobranca.write - Emissão e cancelamento de boletos
pagamento-boleto.write - Pagamento de titulo com código de barras
pagamento-boleto.read - Obter dados completos do titulo a partir do código de barras ou da linha digitável
pagamento-darf.write - Pagamento de DARF sem código de barras
cob.write - Emissão / alteração de pix cobrança imediata
cob.read - Consulta de pix cobrança imediata
cobv.write - Emissão / alteração de pix cobrança com vencimento
cobv.read - Consulta de cobrança com vencimento
pix.write - Solicitação de devolução de pix
pix.read - Consulta de pix
webhook.read - Consulta do webhook
webhook.write - Alteração do webhook
payloadlocation.write - Criação de location do payload
payloadlocation.read - Consulta de locations de payloads
pagamento-pix.write - Pagamento de pix
pagamento-pix.read - Consulta de pix
webhook-banking.write - Alteração de webhooks da API Banking
webhook-banking.read - Consulta do webhooks da API Banking
*/
func (c *Credentials) SetScopesFromString(scopes string) *Credentials {
	if scopes == "" {
		return c
	}

	scopesStrings := strings.Split(scopes, ",")
	c.scopes = make([]Scope, len(scopesStrings))
	for i, scope := range scopesStrings {
		c.scopes[i] = Scope(scope)
	}
	return c
}
