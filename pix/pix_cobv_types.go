package pix

type ModalidadeMulta int

const (
	ModalidadeMultaValorFixo  ModalidadeMulta = 1
	ModalidadeMultaPercentual ModalidadeMulta = 2
)

type ModalidadeJuros int

const (
	ModalidadeJurosValorDiasCorridos  ModalidadeJuros = 1
	ModalidadeJurosPercentualDia      ModalidadeJuros = 2
	ModalidadeJurosPercentualMes      ModalidadeJuros = 3
	ModalidadeJurosPercentualAno      ModalidadeJuros = 4
	ModalidadeJurosValorDiasUteis     ModalidadeJuros = 5
	ModalidadeJurosPercentualDiaUteis ModalidadeJuros = 6
	ModalidadeJurosPercentualMesUteis ModalidadeJuros = 7
	ModalidadeJurosPercentualAnoUteis ModalidadeJuros = 8
)

type ModalidadeAbatimento int

const (
	ModalidadeAbatimentoValorFixo  ModalidadeAbatimento = 1
	ModalidadeAbatimentoPercentual ModalidadeAbatimento = 2
)

type ModalidadeDesconto string

const (
	ModalidadeDescontoValorFixoAteDataInformada          ModalidadeDesconto = "1"
	ModalidadeDescontoPercentualAteDataInformada         ModalidadeDesconto = "2"
	ModalidadeDescontoValorPorAntecipacaoDiaCorrido      ModalidadeDesconto = "3"
	ModalidadeDescontoValorPorAntecipacaoDiaUtil         ModalidadeDesconto = "4"
	ModalidadeDescontoPercentualPorAntecipacaoDiaCorrido ModalidadeDesconto = "5"
	ModalidadeDescontoPercentualPorAntecipacaoDiaUtil    ModalidadeDesconto = "6"
)

// Calendário com vencimento
type CalendarioComVencimento struct {
	Criacao                string `json:"criacao"`                          // Data e hora de criação da cobrança
	DataDeVencimento       string `json:"dataDeVencimento"`                 // Data de vencimento da cobrança
	ValidadeAposVencimento int32  `json:"validadeAposVencimento,omitempty"` // Quantidade de dias após o vencimento que a cobrança poderá ser paga antes de ser considerada expirada.
}

// Componente de valor de multa
type ComponenteValorMulta struct {
	Modalidade ModalidadeMulta `json:"modalidade"` // Modalidade de multa
	ValorPerc  string          `json:"valorPerc"`  // Multa aplicada ao documento, em valor absoluto ou percentual do valor original do documento.
}

// Componente de valor de juros
type ComponenteValorJuros struct {
	Modalidade ModalidadeJuros `json:"modalidade"` // Modalidade de juros
	ValorPerc  string          `json:"valorPerc"`  // Juros aplicados ao documento, em valor absoluto ou percentual do valor original do documento.
}

// Componente de valor de abatimento
type ComponenteValorAbatimento struct {
	Modalidade ModalidadeAbatimento `json:"modalidade"` // Modalidade de abatimentos
	ValorPerc  string               `json:"valorPerc"`  // Abatimentos ou outras deduções aplicadas ao documento, em valor absoluto ou percentual do valor original do documento.
}

// Descontos por pagamento antecipado, com data fixa.
type DescontoDataFixa struct {
	// Descontos por pagamento antecipado, com data fixa.
	// Matriz com até três elementos, sendo que cada elemento é composto por um par "data e valorPerc", para estabelecer descontos percentuais ou absolutos, até aquela data de pagamento.
	// Trata-se de uma data, no formato YYYY-MM-DD, segundo ISO 8601.
	// A data de desconto obrigatoriamente deverá ser menor que a data de vencimento da cobrança.
	Data string `json:"data"`

	// Desconto em valor absoluto ou percentual por dia, útil ou corrido, conforme valor.desconto.modalidade
	ValorPerc string `json:"valorPerc"`
}

// Componente de valor de desconto
type ComponenteValorDesconto struct {
	Modalidade       ModalidadeDesconto  `json:"modalidade"`                 // Modalidade de descontos
	ValorPerc        string              `json:"valorPerc,omitempty"`        // Abatimentos ou outras deduções aplicadas ao documento, em valor absoluto ou percentual do valor original do documento.
	DescontoDataFixa []*DescontoDataFixa `json:"descontoDataFixa,omitempty"` // Descontos absolutos aplicados à cobrança.
}

// Valor da cobrança com vencimento
type ValorCobrancaComVencimento struct {
	Original string `json:"original"` // Valor original da cobrança

	Multa *ComponenteValorMulta `json:"multa,omitempty"` // Multa aplicada à cobrança

	Juros *ComponenteValorJuros `json:"juros,omitempty"` // Juro aplicado à cobrança

	Abatimento *ComponenteValorAbatimento `json:"abatimento,omitempty"` // Abatimento aplicado à cobrança

	Desconto *ComponenteValorDesconto `json:"desconto,omitempty"` // Desconto aplicado à cobrança
}

// CobrancaComVencimentoRequest - Cobrança com vencimento
type CobrancaComVencimentoRequest struct {
	Calendario         *CalendarioComVencimento    `json:"calendario"`               // Calendário com vencimento
	Devedor            *Identificador              `json:"devedor"`                  // Devedor
	Loc                *Loc                        `json:"loc,omitempty"`            // Identificador da localização do payload
	Valor              *ValorCobrancaComVencimento `json:"valor"`                    // Valor da cobrança
	Chave              string                      `json:"chave"`                    // Chave Pix do recebedor
	InfoAdicionais     []*InfoAdicional            `json:"infoAdicionais,omitempty"` // Informações adicionais
	SolicitacaoPagador string                      `json:"solicitacaoPagador"`       // Solicitação do pagador
}

// CobrancaComVencimentoResponse - Cobrança com vencimento
type CobrancaComVencimentoResponse struct {
	TxId               string                      `json:"txid"`                     // Identificador da transação
	Calendario         *CalendarioComVencimento    `json:"calendario,omitempty"`     // Calendário com vencimento
	Devedor            *Identificador              `json:"devedor"`                  // Devedor
	Recebedor          *Identificador              `json:"recebedor"`                // Recebedor
	Loc                *Loc                        `json:"loc,omitempty"`            // Identificador da localização do payload
	Status             CobrancaStatus              `json:"status"`                   // Status da cobrança
	Chave              string                      `json:"chave"`                    // Chave Pix do recebedor
	Valor              *ValorCobrancaComVencimento `json:"valor"`                    // Valor da cobrança
	PixCopiaECola      string                      `json:"pixCopiaECola,omitempty"`  // Este campo retorna o valor do Pix Copia e Cola correspondente à cobrança. Trata-se da sequência de caracteres que representa o BR Code.
	InfoAdicionais     []*InfoAdicional            `json:"infoAdicionais,omitempty"` // Informações adicionais
	SolicitacaoPagador string                      `json:"solicitacaoPagador"`       // Solicitação do pagador
	Revisao            int32                       `json:"revisao,omitempty"`        // Revisão
	Pix                []interface{}               `json:"pix,omitempty"`            // Pix
}

// ConsultarCobrancasComVencimentoRequest representa a requisição para consultar cobranças imediatas.
type ConsultarCobrancasComVencimentoRequest struct {
	Inicio                  string         `json:"inicio,omitempty" time_format:"2006-01-02"` // Data de início da consulta
	Fim                     string         `json:"fim,omitempty" time_format:"2006-01-02"`    // Data de fim da consulta
	Cpf                     string         `json:"cpf,omitempty"`                             // CPF do pagador
	Cnpj                    string         `json:"cnpj,omitempty"`                            // CNPJ do pagador
	LocationPresente        bool           `json:"locationPresente,omitempty"`                // Indica se a localização do payload está presente
	PaginacaoPaginaAtual    int32          `json:"paginacao.paginaAtual,omitempty"`           // Página atual
	PaginacaoItensPorPagina int32          `json:"paginacao.itensPorPagina,omitempty"`        // Itens por página
	Status                  CobrancaStatus `json:"status,omitempty"`                          // Status da cobrança
	LoteCobVId              string         `json:"loteCobVId,omitempty"`                      // Identificador do lote de cobranças
}

// ConsultarCobrancasComVencimentoResponse - Consulta cobranças imediatas com vencimento
type ConsultarCobrancasComVencimentoResponse struct {
	Parametros *ParametrosConsulta              `json:"parametros,omitempty"` // Parâmetros de consulta
	Cobs       []*CobrancaComVencimentoResponse `json:"cobs,omitempty"`       // Cobranças imediatas
}
