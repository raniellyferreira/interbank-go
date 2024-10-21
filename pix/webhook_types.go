package pix

// WebhookResponse representa a resposta da criação de um webhook
type WebhookResponse struct {
	WebhookURL string `json:"webhookUrl"`
}

// Callback representa um evento de webhook
type Callback struct {
	WebhookUrl      string `json:"webhookUrl"`
	NumeroTentativa int32  `json:"numeroTentativa"`
	DataHoraDisparo string `json:"dataHoraDisparo"`
	Sucesso         bool   `json:"sucesso"`
	HttpStatus      int32  `json:"httpStatus"`
	MensagemErro    string `json:"mensagemErro"`
}

type CallbackResponse struct {
	TotalElementos int64       `json:"totalElementos"`
	TotalPaginas   int64       `json:"totalPaginas"`
	PrimeiraPagina bool        `json:"primeiraPagina"`
	UltimaPagina   bool        `json:"ultimaPagina"`
	Data           []*Callback `json:"data"`
}
