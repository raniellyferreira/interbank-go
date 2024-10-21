package pix

type ModalidadeAgente string

const (
	ModalidadeAgenteEstabelecimentoComercial                           ModalidadeAgente = "AGTEC"
	ModalidadeAgenteOutraEspecieDePessoaJuridicaOuCorrespondenteNoPais ModalidadeAgente = "AGTOT"
	ModalidadeAgenteFacilitadorDeServicoDeSaque                        ModalidadeAgente = "AGPSS"
)

type ModalidadeAlteracao int

const (
	ModalidadeAlteracaoNaoPermitido ModalidadeAlteracao = 0
	ModalidadeAlteracaoPermitido    ModalidadeAlteracao = 1
)

type CalendarioSemVencimento struct {
	Expiracao int32 `json:"expiracao"` // Tempo de expiração em segundos.
}

type Identificador struct {
	Nome         string `json:"nome"`
	Cpf          string `json:"cpf,omitempty"`
	Cnpj         string `json:"cnpj,omitempty"`
	NomeFantasia string `json:"nomeFantasia,omitempty"`
	Cidade       string `json:"cidade,omitempty"`
	Uf           string `json:"uf,omitempty"`
	Cep          string `json:"cep,omitempty"`
	Logradouro   string `json:"logradouro,omitempty"`
}

type ValorCobranca struct {
	Original            string              `json:"original"`                      // Valor original da cobrança
	ModalidadeAlteracao ModalidadeAlteracao `json:"modalidadeAlteracao,omitempty"` // Modalidade de alteração

	Retirada *ComponenteValorPix `json:"retirada,omitempty"` // Componente de valor de retirada

	Saque *ComponenteValorPix `json:"saque,omitempty"` // Componente de valor de saque

	Troco *ComponenteValorPix `json:"troco,omitempty"` // Componente de valor de troco
}

type ComponenteValorPix struct {
	Valor                     string              `json:"valor,omitempty"`                     // Valor da cobrança
	ModalidadeAgente          ModalidadeAgente    `json:"modalidadeAgente,omitempty"`          // Modalidade de agente
	ModalidadeAlteracao       ModalidadeAlteracao `json:"modalidadeAlteracao,omitempty"`       // Modalidade de alteração
	PrestadorDoServicoDeSaque string              `json:"prestadorDoServicoDeSaque,omitempty"` // Facilitador de Serviço de Saque
}

type InfoAdicional struct {
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
}

type CobrancaImediataRequest struct {
	Calendario         *CalendarioSemVencimento `json:"calendario"`               // Expiração
	Devedor            *Identificador           `json:"devedor"`                  // Devedor
	Valor              *ValorCobranca           `json:"valor"`                    // Valor da cobrança
	Chave              string                   `json:"chave"`                    // Chave Pix do recebedor
	InfoAdicionais     []*InfoAdicional         `json:"infoAdicionais,omitempty"` // Cada respectiva informação adicional contida na lista (nome e valor) deve ser apresentada ao pagador.
	SolicitacaoPagador string                   `json:"solicitacaoPagador"`       // O campo solicitacaoPagador determina um texto a ser apresentado ao pagador para que ele possa digitar uma informação correlata, em formato livre, a ser enviada ao recebedor. Esse texto está limitado a 140 caracteres.

	Loc    *Loc           `json:"loc,omitempty"`    // Identificador da localização do payload
	Status CobrancaStatus `json:"status,omitempty"` // Status da Cobrança
}

type CalendarioResponse struct {
	Criacao   string `json:"criacao,omitempty"`
	Expiracao int32  `json:"expiracao,omitempty"`
}

type TipoCobranca string

const (
	CobrancaSemValidade TipoCobranca = "cob"
	CobrancaComValidade TipoCobranca = "cobv"
)

type CobrancaStatus string

const (
	CobrancaStatusAtiva                        CobrancaStatus = "ATIVA"
	CobrancaStatusConcluida                    CobrancaStatus = "CONCLUIDA"
	CobrancaStatusRemovidaPeloUsuarioRecebedor CobrancaStatus = "REMOVIDA_PELO_USUARIO_RECEBEDOR"
	CobrancaStatusRemovidaPeloPSP              CobrancaStatus = "REMOVIDA_PELO_PSP"
)

// Identificador da localização do payload.
type Loc struct {
	ID       int64        `json:"id"`                 // Identificador da location a ser informada na criação da cobrança.
	TipoCob  TipoCobranca `json:"tipoCob"`            // Tipo da cobrança - Enum: "cob" "cobv"
	Location string       `json:"location,omitempty"` // Localização do payload
	Criacao  string       `json:"criacao,omitempty"`  // Data e hora de criação da location
}

type CobrancaImediataResponse struct {
	Calendario     *CalendarioResponse `json:"calendario"`     // Expiração
	TxId           string              `json:"txid"`           // Identificador da transação
	Loc            *Loc                `json:"loc"`            // Identificador da localização do payload
	Location       string              `json:"location"`       // Localização do Payload a ser informada na criação da cobrança.
	Status         CobrancaStatus      `json:"status"`         // Status da Cobrança
	Devedor        *Identificador      `json:"devedor"`        // Devedor
	Valor          *ValorCobranca      `json:"valor"`          // Valor da cobrança
	Chave          string              `json:"chave"`          // Chave Pix do recebedor
	InfoAdicionais []*InfoAdicional    `json:"infoAdicionais"` // 	Cada respectiva informação adicional contida na lista (nome e valor) deve ser apresentada ao pagador.

	SolicitacaoPagador string `json:"solicitacaoPagador"` // O campo solicitacaoPagador determina um texto a ser apresentado ao pagador para que ele possa digitar uma informação correlata, em formato livre, a ser enviada ao recebedor. Esse texto está limitado a 140 caracteres.

	PixCopiaECola string `json:"pixCopiaECola,omitempty"` // Este campo retorna o valor do Pix Copia e Cola correspondente à cobrança. Trata-se da sequência de caracteres que representa o BR Code.

	/*
		Denota a revisão da cobrança. Sempre começa em zero. Sempre varia em acréscimos de 1.

		O incremento em uma cobrança deve ocorrer sempre que um objeto da cobrança em questão for alterado. O campo loc é uma exceção a esta regra.

		Se em uma determinada alteração em uma cobrança, o único campo alterado for o campo loc, então esta operação não incrementa a revisão da cobrança.

		O campo loc não ocasiona uma alteração na cobrança em si. Não é necessário armazenar histórico das alterações do campo loc para uma determinada cobrança. Para os outros campos da cobrança, registra-se histórico.
	*/
	Revisao int32 `json:"revisao"`
}

// ConsultarCobrancasImediatasRequest representa a requisição para consultar cobranças imediatas.
type ConsultarCobrancasImediatasRequest struct {
	Inicio                  string         `json:"inicio,omitempty"`                   // Data de início da consulta
	Fim                     string         `json:"fim,omitempty"`                      // Data de fim da consulta
	Cpf                     string         `json:"cpf,omitempty"`                      // CPF do pagador
	Cnpj                    string         `json:"cnpj,omitempty"`                     // CNPJ do pagador
	LocationPresente        bool           `json:"locationPresente,omitempty"`         // Indica se a localização do payload está presente
	PaginacaoPaginaAtual    int32          `json:"paginacao.paginaAtual,omitempty"`    // Página atual
	PaginacaoItensPorPagina int32          `json:"paginacao.itensPorPagina,omitempty"` // Itens por página
	Status                  CobrancaStatus `json:"status,omitempty"`                   // Status da cobrança
}

// Paginacao representa a paginação.
type Paginacao struct {
	PaginaAtual            int32 `json:"paginaAtual,omitempty"`            // Página atual
	ItensPorPagina         int32 `json:"itensPorPagina,omitempty"`         // Itens por página
	QuantidadeDePaginas    int32 `json:"quantidadeDePaginas,omitempty"`    // Quantidade de páginas
	QuantidadeTotalDeItens int64 `json:"quantidadeTotalDeItens,omitempty"` // Quantidade total de itens
}

// ParametrosConsulta representa os parâmetros da consulta de cobranças imediatas.
type ParametrosConsulta struct {
	Inicio    string     `json:"inicio,omitempty"`    // Data de início da consulta
	Fim       string     `json:"fim,omitempty"`       // Data de fim da consulta
	Paginacao *Paginacao `json:"paginacao,omitempty"` // Paginação
}

// ConsultarCobrancasImediatasResponse representa a resposta da consulta de cobranças imediatas.
type ConsultarCobrancasImediatasResponse struct {
	// Parametros representa os parâmetros da consulta de cobranças imediatas.
	Parametros *ParametrosConsulta `json:"parametros"`

	// Cobranças imediatas.
	Cobs []*CobrancaImediataResponse `json:"cobs"`
}
