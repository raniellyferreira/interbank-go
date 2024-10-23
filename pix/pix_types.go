package pix

import (
	"time"

	"github.com/google/uuid"
)

// RecebidosRequest representa a requisição de pix recebidos
type RecebidosRequest struct {
	Inicio                  string `json:"inicio"`                             // Data de início
	Fim                     string `json:"fim"`                                // Data de fim
	TxID                    string `json:"txId,omitempty"`                     // Identificador da transação
	TxIDPresente            bool   `json:"txIdPresente,omitempty"`             // Identificador da transação presente
	DevolucaoPresente       bool   `json:"devolucaoPresente,omitempty"`        // Devolução presente
	CPF                     string `json:"cpf,omitempty"`                      // CPF
	CNPJ                    string `json:"cnpj,omitempty"`                     // CNPJ
	PaginacaoPaginaAtual    int    `json:"paginacao.paginaAtual,omitempty"`    // Página atual
	PaginacaoItensPorPagina int    `json:"paginacao.itensPorPagina,omitempty"` // Itens por página
}

// RecebidosResponse representa a resposta de pix recebidos
type RecebidosResponse struct {
	Parametros *ParametrosConsulta `json:"parametros"`
	Pix        []*Pix              `json:"pix"`
}

// Pix representa a resposta de um pix
type Pix struct {
	// EndToEndID é o identificador único do pix
	EndToEndID string `json:"endToEndId"`

	// Txid é o identificador da transação
	Txid string `json:"txid"`

	// Valor é o valor do pix
	Valor string `json:"valor"`

	/*
		Campo chave do recebedor conforme atribuído na respectiva PACS008.
		Os tipos de chave podem ser: telefone, e-mail, cpf/cnpj ou EVP.
		O formato das chaves pode ser encontrado na seção "Formatação das chaves do DICT no BR Code" do Manual de Padrões para iniciação do Pix.
	*/
	Chave string `json:"chave"`

	/*
		O objetivo dessa estrutura é explicar os elementos de composição do valor do Pix, incluindo informações sobre as multas, juros, descontos e abatimentos quando o Pix for relativo a cobranças com vencimento.

		Regras da estrutura:

		- O valor do Pix é igual a: (original.valor + saque.valor + troco.valor) + multa.valor + juros.valor – abatimento.valor – desconto.valor considerando-se apenas os campos que estiverem presentes para cada tipo de cobrança pago.
		- As estruturas saque e troco só serão retornadas quando o Pix for relativo a um Pix Saque ou Pix Troco, respectivamente, e as demais estruturas (juros, multa, abatimento e desconto) só serão pertinentes aos Pix de pagamentos das cobranças com vencimento.
		- Não pode haver simultaneamente uma subsestrutura do tipo saque e outra do tipo troco;
		- Não há restrição na ordem das subestruturas.
	*/
	ComponentesValor *ComponentesValor `json:"componentesValor,omitempty"`

	// Horário em que o Pix foi processado no PSP.
	Horario string `json:"horario"`

	// Informação livre do pagador
	InfoPagador string `json:"infoPagador"`

	// Devoluções
	Devolucoes []*DevolucaoResponse `json:"devolucoes"`
}

// ComponentesValor representa os componentes do valor de um pix
type ComponentesValor struct {
	// Original é o valor original do pix
	Original *ComponenteValor `json:"original,omitempty"`

	// Saque é o valor do saque do pix
	Saque *ComponenteValor `json:"saque,omitempty"`

	// Troco é o valor do troco do pix
	Troco *ComponenteValor `json:"troco,omitempty"`

	// Multa é o valor da multa do pix
	Multa *ComponenteValor `json:"multa,omitempty"`

	// Juros é o valor dos juros do pix
	Juros *ComponenteValor `json:"juros,omitempty"`

	// Desconto é o valor do desconto do pix
	Desconto *ComponenteValor `json:"desconto,omitempty"`

	// Abatimento é o valor do abatimento do pix
	Abatimento *ComponenteValor `json:"abatimento,omitempty"`
}

// ComponenteValor representa o valor original de um pix
type ComponenteValor struct {
	// Valor é o valor original do pix
	Valor float64 `json:"valor,omitempty"`

	ModalidadeAgente ModalidadeAgente `json:"modalidadeAgente"`

	PrestadorDeServicoDeSaque string `json:"prestadorDeServicoDeSaque"`
}

// Horario representa o horário de um pix
type Horario struct {
	// Solicitacao é o horário da solicitação do pix
	Solicitacao time.Time `json:"solicitacao"`
}

type DevolucaoStatus string

const (
	// DevolucaoStatusEmProcessamento representa o status de devolução em processamento
	DevolucaoStatusEmProcessamento DevolucaoStatus = "EM_PROCESSAMENTO"

	// DevolucaoStatusDevolvido representa o status de devolução devolvido
	DevolucaoStatusDevolvido DevolucaoStatus = "DEVOLVIDO"

	// DevolucaoStatusNaoRealizado representa o status de devolução não realizado
	DevolucaoStatusNaoRealizado DevolucaoStatus = "NAO_REALIZADO"
)

// DevolucaoResponse representa a devolução de um pix
type DevolucaoResponse struct {
	// ID é o identificador da devolução
	ID string `json:"id"`

	// RtrID é o identificador do RTR
	RtrID string `json:"rtrId"`

	// Valor é o valor da devolução
	Valor float64 `json:"valor"`

	// Horario é o horário da devolução
	Horario *Horario `json:"horario"`

	// Status é o status da devolução
	Status DevolucaoStatus `json:"status"`

	// Descrição é a descrição da devolução
	Descricao string `json:"descricao,omitempty"`
}

type NaturezaDevolucaoPix string

const (
	// NaturezaDevolucaoOriginal é a natureza original
	NaturezaDevolucaoOriginal NaturezaDevolucaoPix = "ORIGINAL"

	// NaturezaDevolucaoRetirada é a natureza retirada
	NaturezaDevolucaoRetirada NaturezaDevolucaoPix = "RETIRADA"
)

type SolicitarDevolucaoPixRequest struct {
	// LocalUniqId é o identificador único da devolução
	LocalUniqId string `json:"-"`

	// EndToEndID é o identificador único do pix
	EndToEndID string `json:"-"`

	// Valor é o valor da devolução
	Valor float64 `json:"valor"`

	/*
		Default: "ORIGINAL"
		Enum: "ORIGINAL" "RETIRADA"
		Indica qual é a natureza da devolução solicitada. Uma solicitação de devolução pelo usuário recebedor pode ser relacionada a um Pix comum (com código: MD06 da pacs.004), ou a um Pix de Saque ou Troco (com códigos possíveis: MD06 e SL02 da pacs.004). Na ausência deste campo a natureza deve ser interpretada como sendo de um Pix comum (ORIGINAL).

		As naturezas são assim definidas:

		ORIGINAL: quando a devolução é solicitada pelo usuário recebedor e se refere a um Pix comum ou ao valor da compra em um Pix Troco (MD06);
		RETIRADA: quando a devolução é solicitada pelo usuário recebedor e se refere a um Pix Saque ou ao valor do troco em um Pix Troco (SL02).
		Os valores de devoluções são sempre limitados aos valores máximos a seguir:

		Pix comum: o valor da devolução é limitado ao valor do próprio Pix (a natureza nesse caso deve ser: ORIGINAL);
		Pix Saque: o valor da devolução é limitado ao valor da retirada (a natureza nesse caso deve ser: RETIRADA); e
		Pix Troco: o valor da devolução é limitado ao valor relativo à compra ou ao troco:
		Quando a devolução for referente à compra, o valor limita-se ao valor da compra (a natureza nesse caso deve ser ORIGINAL); e
		Quando a devolução for referente ao troco, o valor limita-se ao valor do troco (a natureza nesse caso deve ser RETIRADA).
	*/
	Natureza NaturezaDevolucaoPix `json:"natureza"`

	// Descrição é a descrição da devolução
	Descricao string `json:"descricao,omitempty"`
}

func (r *SolicitarDevolucaoPixRequest) GetLocalUniqId() string {
	if r.LocalUniqId == "" {
		r.LocalUniqId = uuid.NewString()
	}
	return r.LocalUniqId
}
