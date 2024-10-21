package cobranca

type TipoPessoa string

const (
	// PessoaFisica represents a physical person
	PessoaFisica TipoPessoa = "FISICA"

	// PessoaJuridica represents a legal person
	PessoaJuridica TipoPessoa = "JURIDICA"
)

type Pessoa struct {
	Nome       string     `json:"nome"`       // Nome do pagador
	CpfCnpj    string     `json:"cpfCnpj"`    // CPF ou CNPJ do pagador
	TipoPessoa TipoPessoa `json:"tipoPessoa"` // Tipo de pessoa do pagador

	// Campos opcionais
	Email       string `json:"email,omitempty"`       // Email do pagador
	Ddd         string `json:"ddd,omitempty"`         // DDD do telefone do pagador
	Telefone    string `json:"telefone,omitempty"`    // Telefone do pagador
	Numero      string `json:"numero,omitempty"`      // Número do telefone do pagador
	Complemento string `json:"complemento,omitempty"` // Complemento do endereço do pagador
	Endereco    string `json:"endereco,omitempty"`    // Endereço do pagador
	Bairro      string `json:"bairro,omitempty"`      // Bairro do pagador
	Cidade      string `json:"cidade,omitempty"`      // Cidade do pagador
	Uf          string `json:"uf,omitempty"`          // UF do pagador
	Cep         string `json:"cep,omitempty"`         // CEP do pagador
}

type FormaRecebimento string

const (
	// FormasRecebimentoCobrancaPix representa a forma de recebimento pix
	FormasRecebimentoCobrancaPix FormaRecebimento = "PIX"

	// FormasRecebimentoCobrancaBoleto representa a forma de recebimento boleto
	FormasRecebimentoCobrancaBoleto FormaRecebimento = "BOLETO"
)

type ComponenteValor struct {
	Codigo         string `json:"codigo"`                   // Código da condição de pagamento
	Taxa           string `json:"taxa,omitempty"`           // Taxa da condição de pagamento
	Valor          int    `json:"valor,omitempty"`          // Valor da condição de pagamento
	QuantidadeDias int    `json:"quantidadeDias,omitempty"` // Quantidade de dias da condição de pagamento
}

type CobrancaMensagem struct {
	Linha1 string `json:"linha1,omitempty"`
	Linha2 string `json:"linha2,omitempty"`
	Linha3 string `json:"linha3,omitempty"`
	Linha4 string `json:"linha4,omitempty"`
	Linha5 string `json:"linha5,omitempty"`
}

type EmitirRequest struct {
	SeuNumero         string            `json:"seuNumero"`         // Campo Seu Número do título
	ValorNominal      string            `json:"valorNominal"`      // Valor nominal do título
	DataVencimento    string            `json:"dataVencimento"`    // Data de vencimento do título
	NumDiasAgenda     string            `json:"numDiasAgenda"`     // Número de dias corridos após o vencimento para o cancelamento efetivo automático da cobrança. (de 0 até 60)
	FormasRecebimento string            `json:"formasRecebimento"` // Lista com as formas de recebimento de uma cobrança, separadas por vírgula.
	Pagador           *Pessoa           `json:"pagador"`           // Dados do pagador
	BeneficiarioFinal *Pessoa           `json:"beneficiarioFinal"` // Dados do beneficiário final
	Desconto          *ComponenteValor  `json:"desconto"`          // Dados do desconto
	Multa             *ComponenteValor  `json:"multa"`             // Dados da multa
	Mora              *ComponenteValor  `json:"mora"`              // Dados da mora
	Mensagem          *CobrancaMensagem `json:"mensagem"`          // Mensagem da cobrança
}

type EmitirResponse struct {
	CodigoSolicitacao string `json:"codigoSolicitacao"`
}
