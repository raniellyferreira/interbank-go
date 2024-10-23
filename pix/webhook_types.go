package pix

// WebhookCall representa um evento de webhook para o endpoint do callback
type WebhookCall struct {
	Pix []*Pix `json:"pix"`
}

// WebhookResponse representa a resposta da criação de um webhook
type WebhookResponse struct {
	WebhookURL string `json:"webhookUrl"`
}

// CallbackItem representa um evento de webhook
type CallbackItem struct {
	WebhookUrl      string `json:"webhookUrl"`
	NumeroTentativa int32  `json:"numeroTentativa"`
	DataHoraDisparo string `json:"dataHoraDisparo"`
	Sucesso         bool   `json:"sucesso"`
	HttpStatus      int32  `json:"httpStatus"`
	MensagemErro    string `json:"mensagemErro"`
}

type CallbacksResponse struct {
	TotalElementos int64           `json:"totalElementos"`
	TotalPaginas   int64           `json:"totalPaginas"`
	PrimeiraPagina bool            `json:"primeiraPagina"`
	UltimaPagina   bool            `json:"ultimaPagina"`
	Data           []*CallbackItem `json:"data"`
}

// ConsultarWebhooksCallbacksRequest representa a requisição para consultar os eventos de um webhook
type ConsultarWebhooksCallbacksRequest struct {
	TxID           string `json:"txid,omitempty"`
	DataHoraInicio string `json:"dataHoraInicio"`
	DataHoraFim    string `json:"dataHoraFim"`
	Pagina         int32  `json:"pagina,omitempty"`
	TamanhoPagina  int32  `json:"tamanhoPagina,omitempty"`
}
