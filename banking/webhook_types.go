package banking

type TipoWebhook string

const (
	TipoWebhookPagamentoPix    TipoWebhook = "pix-pagamento"
	TipoWebhookPagamentoBoleto TipoWebhook = "boleto-pagamento"
)

// WebhookResponse é a resposta de uma requisição de consulta de webhook
type WebhookResponse struct {
	WebhookUrl string `json:"webhookUrl"` // URL de configuração do webhook. Deve iniciar obrigatoriamente com https://
	Criacao    string `json:"criacao"`    // Data e hora em que o webhook foi cadastrado. Representada de acordo com a RFC3339
}

// WebhookCallbacksRequest é a requisição para consultar os eventos de um webhook
type WebhookCallbacksRequest struct {
	DataHoraInicio  string `json:"dataHoraInicio"`            // Formato aceito: yyyy-MM-dd'T'HH:mm[:ss][.SSS]XXX
	DataHoraFim     string `json:"dataHoraFim"`               // Formato aceito: yyyy-MM-dd'T'HH:mm[:ss][.SSS]XXX
	Pagina          int32  `json:"pagina,omitempty"`          // Posição da página na lista de dados
	TamanhoPagina   int32  `json:"tamanhoPagina,omitempty"`   // Tamanho da página
	EndToEnd        string `json:"endToEnd,omitempty"`        // EndToEnd do callback, caso queira filtrar as notificações de algum Pix Pagamento específico. Exclusivo para o PATH PARAMETERS pix-pagamento.
	CodigoTransacao string `json:"codigoTransacao,omitempty"` // CodigoTransacao do callback, caso queira filtrar as notificações de algum Pagamento Boleto específico. Exclusivo para o PATH PARAMETERS boleto-pagamento.
}

// WebhookCall é um evento de um webhook
type WebhookCall struct {
	WebhookUrl      string `json:"webhookUrl"`      // URL do webhook
	NumeroTentativa int32  `json:"numeroTentativa"` // Número da tentativa
	DataEnvio       string `json:"dataEnvio"`       // Data e hora de envio
	Sucesso         bool   `json:"sucesso"`         // Indica se a tentativa foi bem sucedida
	HttpStatus      int32  `json:"httpStatus"`      // Status HTTP da tentativa
	MensagemErro    string `json:"mensagemErro"`    // Mensagem de erro
}

// WebhookCallbacksResponse é a resposta de uma requisição de consulta de callbacks
type WebhookCallbacksResponse struct {
	TotalElementos int64          `json:"totalElementos"` // Total de elementos
	TotalPaginas   int32          `json:"totalPaginas"`   // Total de páginas
	PrimeiraPagina bool           `json:"primeiraPagina"` // Indica se é a primeira página
	UltimaPagina   bool           `json:"ultimaPagina"`   // Indica se é a última página
	Data           []*WebhookCall `json:"data"`           // Lista de callbacks
}
