package banking

import (
	"encoding/json"

	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// TipoOperacao represents the type of operation
type TipoOperacao string

const (
	// TipoOperacaoDebito represents a debit operation
	TipoOperacaoDebito TipoOperacao = "D"
	// TipoOperacaoCredito represents a credit operation
	TipoOperacaoCredito TipoOperacao = "C"
)

// TipoTransacao represents the type of transaction
type TipoTransacao string

const (
	// TipoTransacaoDebitoEmConta represents a debit operation
	TipoTransacaoDebitoEmConta TipoTransacao = "DEBITO_EM_CONTA"
	// TipoTransacaoDepositoBoleto represents a deposit operation
	TipoTransacaoDepositoBoleto TipoTransacao = "DEPOSITO_BOLETO"
	// TipoTransacaoAntecipacaoRecebiveis represents an anticipation of receivables operation
	TipoTransacaoAntecipacaoRecebiveis TipoTransacao = "ANTECIPACAO_RECEBIVEIS"
	// TipoTransacaoAntecipacaoRecebiveisCartao represents an anticipation of receivables by card operation
	TipoTransacaoAntecipacaoRecebiveisCartao TipoTransacao = "ANTECIPACAO_RECEBIVEIS_CARTAO"
	// TipoTransacaoBoletoCobranca represents a billet collection operation
	TipoTransacaoBoletoCobranca TipoTransacao = "BOLETO_COBRANCA"
	// TipoTransacaoCambio represents a currency exchange operation
	TipoTransacaoCambio TipoTransacao = "CAMBIO"
	// TipoTransacaoCashback represents a cashback operation
	TipoTransacaoCashback TipoTransacao = "CASHBACK"
	// TipoTransacaoCheque represents a check operation
	TipoTransacaoCheque TipoTransacao = "CHEQUE"
	// TipoTransacaoEstorno represents a refund operation
	TipoTransacaoEstorno TipoTransacao = "ESTORNO"
	// TipoTransacaoDomicilioCartao represents a card domicile operation
	TipoTransacaoDomicilioCartao TipoTransacao = "DOMICILIO_CARTAO"
	// TipoTransacaoFinanciamento represents a financing operation
	TipoTransacaoFinanciamento TipoTransacao = "FINANCIAMENTO"
	// TipoTransacaoImposto represents a tax operation
	TipoTransacaoImposto TipoTransacao = "IMPOSTO"
	// TipoTransacaoInterpag represents an interpag operation
	TipoTransacaoInterpag TipoTransacao = "INTERPAG"
	// TipoTransacaoInvestimento represents an investment operation
	TipoTransacaoInvestimento TipoTransacao = "INVESTIMENTO"
	// TipoTransacaoJuros represents an interest operation
	TipoTransacaoJuros TipoTransacao = "JUROS"
	// TipoTransacaoMaquininhaGranito represents a granite machine operation
	TipoTransacaoMaquininhaGranito TipoTransacao = "MAQUININHA_GRANITO"
	// TipoTransacaoMulta represents a fine operation
	TipoTransacaoMulta TipoTransacao = "MULTA"
	// TipoTransacaoOutros represents an other operation
	TipoTransacaoOutros TipoTransacao = "OUTROS"
	// TipoTransacaoPagamento represents a payment operation
	TipoTransacaoPagamento TipoTransacao = "PAGAMENTO"
	// TipoTransacaoPix represents a pix operation
	TipoTransacaoPix TipoTransacao = "PIX"
	// TipoTransacaoProventos represents a profit operation
	TipoTransacaoProventos TipoTransacao = "PROVENTOS"
	// TipoTransacaoSaque represents a withdrawal operation
	TipoTransacaoSaque TipoTransacao = "SAQUE"
	// TipoTransacaoCompraDebito represents a debit purchase operation
	TipoTransacaoCompraDebito TipoTransacao = "COMPRA_DEBITO"
	// TipoTransacaoDebitoAutomatico represents an automatic debit operation
	TipoTransacaoDebitoAutomatico TipoTransacao = "DEBITO_AUTOMATICO"
	// TipoTransacaoTarifa represents a fee operation
	TipoTransacaoTarifa TipoTransacao = "TARIFA"
	// TipoTransacaoTransferencia represents a transfer operation
	TipoTransacaoTransferencia TipoTransacao = "TRANSFERENCIA"
	// TipoTransacaoNIsNotTransaction represents a not transaction operation
	TipoTransacaoIsNotTransaction TipoTransacao = ""
)

// OrigemMovimentacao represents the origin of the movement
type OrigemMovimentacao string

// TipoDetalhe represents the type of detail
type TipoDetalhe string

const (
	// TipoDetalheComplete represents a complete detail
	TipoDetalheComplete TipoDetalhe = "COMPLETE"

	// TipoDetalheIncomplete represents an incomplete detail
	TipoDetalheIncomplete TipoDetalhe = "INCOMPLETE"
)

// Detalhe represents the details of a transaction
type Detalhe interface{}

// DetalhePix represents the details of a pix transaction
type DetalhePix struct {
	TxID                   string             `json:"txId"`                   // ID de identificação do documento para conciliação de pagamentos.
	NomePagador            string             `json:"nomePagador"`            // Nome do pagador.
	DescricaoPix           string             `json:"descricaoPix"`           // Descrição do Pix.
	CpfCnpjPagador         string             `json:"cpfCnpjPagador"`         // CPF ou CNPJ do pagador.
	ContaBancariaRecebedor string             `json:"contaBancariaRecebedor"` // Conta bancária do recebedor.
	NomeEmpresaPagador     string             `json:"nomeEmpresaPagador"`     // Nome da empresa do pagador.
	TipoDetalhe            TipoDetalhe        `json:"tipoDetalhe"`            // Tipo de detalhe.
	EndToEndId             string             `json:"endToEndId"`             // ID único para identificação do pagamento Pix.
	ChavePixRecebedor      string             `json:"chavePixRecebedor"`      // Chave Pix do recebedor.
	NomeEmpresaRecebedor   string             `json:"nomeEmpresaRecebedor"`   // Nome da empresa do recebedor.
	NomeRecebedor          string             `json:"nomeRecebedor"`          // Nome do recebedor.
	AgenciaRecebedor       string             `json:"agenciaRecebedor"`       // Agência do recebedor.
	CpfCnpjRecebedor       string             `json:"cpfCnpjRecebedor"`       // CPF ou CNPJ do recebedor.
	OrigemMovimentacao     OrigemMovimentacao `json:"origemMovimentacao"`     // Origem da movimentação.
	CodigoSolicitacao      string             `json:"codigoSolicitacao"`      // Código de solicitação.
}

// DetalheBoletoCobranca represents the details of a billet collection
type DetalheBoletoCobranca struct {
	DataVencimento string `json:"dataVencimento"` // Data de vencimento do boleto.
	DataTransacao  string `json:"dataTransacao"`  // Data da transação.
	NossoNumero    string `json:"nossoNumero"`    // Nosso número para identificação.
	SeuNumero      string `json:"seuNumero"`      // Seu número para referência.
	CodBarras      string `json:"codBarras"`      // Código de barras do boleto.
	Juros          string `json:"juros"`          // Juros aplicados.
	Multa          string `json:"multa"`          // Multa aplicada.
	Desconto1      string `json:"desconto1"`      // Primeiro desconto oferecido.
	Desconto2      string `json:"desconto2"`      // Segundo desconto oferecido.
	Desconto3      string `json:"desconto3"`      // Terceiro desconto oferecido.
	Nome           string `json:"nome"`           // Nome do pagador.
	DataLimite     string `json:"dataLimite"`     // Data limite para pagamento.
	TipoDetalhe    string `json:"tipoDetalhe"`    // Tipo de detalhe do boleto.
	CpfCnpj        string `json:"cpfCnpj"`        // CPF ou CNPJ do pagador.
	DataEmissao    string `json:"dataEmissao"`    // Data de emissão do boleto.
	Abatimento     string `json:"abatimento"`     // Valor de abatimento aplicado.
}

// DetalheCashback represents the details of a cashback
type DetalheCashback struct {
	ValorCompra string `json:"valorCompra"` // Valor da compra.
	Produto     string `json:"produto"`     // Produto associado ao cashback.
	TipoDetalhe string `json:"tipoDetalhe"` // Tipo de detalhe do cashback.
}

// DetalheCheque represents the details of a check
type DetalheCheque struct {
	Agencia                 string `json:"agencia"`                 // Agência do cheque.
	NumeroChequeBancario    string `json:"numeroChequeBancario"`    // Número do cheque bancário.
	ContaBancaria           string `json:"contaBancaria"`           // Conta bancária associada.
	DataRetorno             string `json:"dataRetorno"`             // Data de retorno do cheque.
	MotivoRetorno           string `json:"motivoRetorno"`           // Motivo do retorno do cheque.
	DescricaoChequeBancario string `json:"descricaoChequeBancario"` // Descrição do cheque bancário.
	NomeEmpresa             string `json:"nomeEmpresa"`             // Nome da empresa associada.
	TipoDetalhe             string `json:"tipoDetalhe"`             // Tipo de detalhe do cheque.
	CodigoAfiliado          string `json:"codigoAfiliado"`          // Código do afiliado.
}

// DetalheCompraDebito represents the details of a debit purchase
type DetalheCompraDebito struct {
	Estabelecimento string `json:"estabelecimento"` // Nome do estabelecimento.
	TipoDetalhe     string `json:"tipoDetalhe"`     // Tipo de detalhe da compra.
}

// DetalheDepositoBoleto represents the details of a billet deposit
type DetalheDepositoBoleto struct {
	DataVencimento string `json:"dataVencimento"` // Data de vencimento do boleto.
	TipoDetalhe    string `json:"tipoDetalhe"`    // Tipo de detalhe do depósito.
	DataEmissao    string `json:"dataEmissao"`    // Data de emissão do boleto.
	NossoNumero    string `json:"nossoNumero"`    // Nosso número para identificação.
	CodBarras      string `json:"codBarras"`      // Código de barras do boleto.
}

// DetalheTransferencia represents the details of a transfer
type DetalheTransferencia struct {
	ContaBancariaPagador   string `json:"contaBancariaPagador"`   // Conta bancária do pagador.
	DescricaoTransferencia string `json:"descricaoTransferencia"` // Descrição da transferência.
	AgenciaPagador         string `json:"agenciaPagador"`         // Agência do pagador.
	BancoRecebedor         string `json:"bancoRecebedor"`         // Banco do recebedor.
	ContaBancariaRecebedor string `json:"contaBancariaRecebedor"` // Conta bancária do recebedor.
	CpfCnpjRecebedor       string `json:"cpfCnpjRecebedor"`       // CPF ou CNPJ do recebedor.
	CpfCnpjPagador         string `json:"cpfCnpjPagador"`         // CPF ou CNPJ do pagador.
	NomePagador            string `json:"nomePagador"`            // Nome do pagador.
	NomeEmpresaPagador     string `json:"nomeEmpresaPagador"`     // Nome da empresa do pagador.
	NomeRecebedor          string `json:"nomeRecebedor"`          // Nome do recebedor.
	TipoDetalhe            string `json:"tipoDetalhe"`            // Tipo de detalhe da transferência.
	IdTransferencia        string `json:"idTransferencia"`        // ID da transferência.
	AgenciaRecebedor       string `json:"agenciaRecebedor"`       // Agência do recebedor.
	DataEfetivacao         string `json:"dataEfetivacao"`         // Data de efetivação da transferência.
}

// DetalhePagamento represents the details of a payment
type DetalhePagamento struct {
	ValorTotal       string `json:"valorTotal"`       // Valor total do pagamento.
	DetalheDescricao string `json:"detalheDescricao"` // Descrição detalhada do pagamento.
	ContaBancaria    string `json:"contaBancaria"`    // Conta bancária associada.
	Agencia          string `json:"agencia"`          // Agência associada.
	Adicionado       string `json:"adicionado"`       // Informação adicionada.
	DataVencimento   string `json:"dataVencimento"`   // Data de vencimento do pagamento.
	CodigoAfiliado   string `json:"codigoAfiliado"`   // Código do afiliado.
	EmpresaEmissora  string `json:"empresaEmissora"`  // Empresa emissora do pagamento.
	ValorOriginal    string `json:"valorOriginal"`    // Valor original do pagamento.
	Desconto         string `json:"desconto"`         // Desconto aplicado.
	CpfCnpj          string `json:"cpfCnpj"`          // CPF ou CNPJ do pagador.
	ValorPrincipal   string `json:"valorPrincipal"`   // Valor principal do pagamento.
	PeriodoApuracao  string `json:"periodoApuracao"`  // Período de apuração.
	ValorAumentado   string `json:"valorAumentado"`   // Valor aumentado.
	CodBarras        string `json:"codBarras"`        // Código de barras associado.
	ValorParcial     string `json:"valorParcial"`     // Valor parcial do pagamento.
	Hora             string `json:"hora"`             // Hora do pagamento.
	Juros            string `json:"juros"`            // Juros aplicados.
	Multa            string `json:"multa"`            // Multa aplicada.
	EmpresaOrigem    string `json:"empresaOrigem"`    // Empresa de origem do pagamento.
	NomeDestinatario string `json:"nomeDestinatario"` // Nome do destinatário do pagamento.
	TipoDetalhe      string `json:"tipoDetalhe"`      // Tipo de detalhe do pagamento.
	NomeOrigem       string `json:"nomeOrigem"`       // Nome de origem do pagamento.
	CodigoReceita    string `json:"codigoReceita"`    // Código de receita associado.
	LinhaDigitavel   string `json:"linhaDigitavel"`   // Linha digitável do pagamento.
	Autenticacao     string `json:"autenticacao"`     // Código de autenticação.
}

// Transacao represents a transaction
type Transacao struct {
	// Common
	DataEntrada   string        `json:"dataEntrada"`   // Data de entrada
	TipoTransacao TipoTransacao `json:"tipoTransacao"` // Tipo de transação
	TipoOperacao  TipoOperacao  `json:"tipoOperacao"`  // Tipo de operação
	Valor         string        `json:"valor"`         // Valor
	Titulo        string        `json:"titulo"`        // Título
	Descricao     string        `json:"descricao"`     // Descrição

	// Only on Extrato
	Cpmf string `json:"cpmf,omitempty"` // CPMF

	// Only on ExtratoCompleto
	IDTransacao   string          `json:"idTransacao,omitempty"`   // ID da transação
	DataInclusao  string          `json:"dataInclusao,omitempty"`  // Data de inclusão
	DataTransacao string          `json:"dataTransacao,omitempty"` // Data da transação
	Detalhes      json.RawMessage `json:"detalhes,omitempty"`      // Detalhes, dependendo do tipo de transação
}

// GetDetalhe unmarshal the Detalhes field into the target
func (t *Transacao) GetDetalhe(target interface{}) error {
	return interutils.JsonUnmarshal(t.Detalhes, target)
}

// ConsultarExtrato response represents the response of the ConsultarExtrato method
type ConsultarExtratoResponse struct {
	// Transacoes represents the transactions
	Transacoes []*Transacao `json:"transacoes"`

	// Only on ExtratoCompleto
	TotalPaginas      int64 `json:"totalPaginas,omitempty"`      // Total de páginas
	TotalElementos    int64 `json:"totalElementos,omitempty"`    // Total de elementos
	UltimaPagina      bool  `json:"ultimaPagina,omitempty"`      // Última página
	PrimeiraPagina    bool  `json:"primeiraPagina,omitempty"`    // Primeira página
	TamanhoPagina     int32 `json:"tamanhoPagina,omitempty"`     // Tamanho da página
	NumeroDeElementos int64 `json:"numeroDeElementos,omitempty"` // Número de elementos
}

// ConsultarExtratoCompletoRequest represents the request of the ConsultarExtratoCompleto method
type ConsultarExtratoCompletoRequest struct {
	DataInicio    string        `json:"dataInicio"`              // Data de início
	DataFim       string        `json:"dataFim"`                 // Data de fim
	Pagina        int32         `json:"pagina,omitempty"`        // Página
	TamanhoPagina int32         `json:"tamanhoPagina,omitempty"` // Tamanho da página
	TipoOperacao  TipoOperacao  `json:"tipoOperacao,omitempty"`  // Tipo de operação
	TipoTransacao TipoTransacao `json:"tipoTransacao,omitempty"` // Tipo de transação
}

// ExportarExtratoResponse represents the response of the ExportarExtrato method
type ExportarExtratoResponse struct {
	Pdf []byte `json:"pdf"` // PDF
}
