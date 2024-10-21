package auth

type Scope string

const (
	ScopeExtratoRead          Scope = "extrato.read"           // Consulta de Extrato e Saldo
	ScopeBoletoCobrancaRead   Scope = "boleto-cobranca.read"   // Consulta de boletos e exportação para PDF
	ScopeBoletoCobrancaWrite  Scope = "boleto-cobranca.write"  // Emissão e cancelamento de boletos
	ScopePagamentoBoletoRead  Scope = "pagamento-boleto.read"  // Obter dados completos do titulo a partir do código de barras ou da linha digitável
	ScopePagamentoBoletoWrite Scope = "pagamento-boleto.write" // Pagamento de titulo com código de barras
	ScopePagamentoDarfWrite   Scope = "pagamento-darf.write"   // Pagamento de DARF sem código de barras
	ScopeCobWrite             Scope = "cob.write"              // Emissão / alteração de pix cobrança imediata
	ScopeCobRead              Scope = "cob.read"               // Consulta de pix cobrança imediata
	ScopeCobVWrite            Scope = "cobv.write"             // Emissão / alteração de pix cobrança com vencimento
	ScopeCobVRead             Scope = "cobv.read"              // Consulta de cobrança com vencimento
	ScopePixWrite             Scope = "pix.write"              // Solicitação de devolução de pix
	ScopePixRead              Scope = "pix.read"               // Consulta de pix
	ScopeWebhookRead          Scope = "webhook.read"           // Consulta do webhook
	ScopeWebhookWrite         Scope = "webhook.write"          // Alteração do webhook
	ScopePayloadLocationWrite Scope = "payloadlocation.write"  // Criação de location do payload
	ScopePayloadLocationRead  Scope = "payloadlocation.read"   // Consulta de locations de payloads
	ScopePagamentoPixWrite    Scope = "pagamento-pix.write"    // Pagamento de pix
	ScopePagamentoPixRead     Scope = "pagamento-pix.read"     // Consulta pagamento de pix
	ScopeWebhookBankingWrite  Scope = "webhook-banking.write"  // Alteração de webhooks da API Banking
	ScopeWebhookBankingRead   Scope = "webhook-banking.read"   // Consulta do webhooks da API Banking
)
