package cobranca

import "github.com/google/uuid"

// CriarWebhookRequest represents a request to create a webhook
type CriarWebhookRequest struct {
	WebhookUrl string `json:"webhookUrl"` // URL to receive the webhook
}

// Webhook represents a webhook
type Webhook struct {
	WebhookUrl  string `json:"webhookUrl"`  // URL to receive the webhook
	Criacao     string `json:"criacao"`     // Creation date
	Atualizacao string `json:"atualizacao"` // Update date
}

// WebhookCallback represents a webhook callback
type WebhookCallback struct {
	WebhookUrl      string `json:"webhookUrl"`             // URL to receive the webhook
	NumeroTentativa int32  `json:"numeroTentativa"`        // Number of attempts
	DataHoraDisparo string `json:"dataHoraDisparo"`        // Dispatch date
	Sucesso         bool   `json:"sucesso"`                // Success
	HttpStatus      int32  `json:"httpStatus"`             // HTTP status
	MensagemErro    string `json:"mensagemErro,omitempty"` // Error message
}

// WebhookCallbacksResponse represents a response to get a webhook callbacks
type WebhookCallbacksResponse struct {
	TotalElementos int64              `json:"totalElementos"` // Total elements
	TotalPaginas   int64              `json:"totalPaginas"`   // Total pages
	PrimeiraPagina bool               `json:"primeiraPagina"` // First page
	UltimaPagina   bool               `json:"ultimaPagina"`   // Last page
	Data           []*WebhookCallback `json:"data"`           // Data
}

// ConsultarWebhookCallbacksRequest represents a request to get a webhook callbacks
type ConsultarWebhookCallbacksRequest struct {
	DataHoraInicio    string    `json:"dataHoraInicio"`              // Start date
	DataHoraFim       string    `json:"dataHoraFim"`                 // End date
	Pagina            int64     `json:"pagina,omitempty"`            // Page
	TamanhoPagina     int32     `json:"tamanhoPagina,omitempty"`     // Page size
	CodigoSolicitacao uuid.UUID `json:"codigoSolicitacao,omitempty"` // Request code
}
