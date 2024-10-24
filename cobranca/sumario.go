package cobranca

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

type FiltrarDataOption string

const (
	FiltrarDataPorVencimento FiltrarDataOption = "VENCIMENTO"
	FiltrarDataPorEmissao    FiltrarDataOption = "EMISSAO"
	FiltrarDataPorPagamento  FiltrarDataOption = "PAGAMENTO"
)

type SituacaoCobranca string

const (
	SituacaoCobrancaRecebido        SituacaoCobranca = "RECEBIDO"
	SituacaoCobrancaAReceber        SituacaoCobranca = "A_RECEBER"
	SituacaoCobrancaMarcadoRecebido SituacaoCobranca = "MARCADO_RECEBIDO"
	SituacaoCobrancaAtrasado        SituacaoCobranca = "ATRASADO"
	SituacaoCobrancaCancelado       SituacaoCobranca = "CANCELADO"
	SituacaoCobrancaExpirado        SituacaoCobranca = "EXPIRADO"
	SituacaoCobrancaFalhaEmissao    SituacaoCobranca = "FALHA_EMISSAO"
	SituacaoCobrancaEmProcessamento SituacaoCobranca = "EM_PROCESSAMENTO"
)

type TipoCobranca string

const (
	TipoCobrancaSimples    TipoCobranca = "SIMPLES"
	TipoCobrancaParcelado  TipoCobranca = "PARCELADO"
	TipoCobrancaRecorrente TipoCobranca = "RECORRENTE"
)

// SumarioRequest representa a requisição de sumário de cobranças
type SumarioRequest struct {
	DataInicial string `json:"dataInicial"`
	DataFinal   string `json:"dataFinal"`

	// Optional fields
	FiltrarDataPor        FiltrarDataOption `json:"filtrarDataPor,omitempty"`
	Situacao              SituacaoCobranca  `json:"situacao,omitempty"`
	TipoCobranca          TipoCobranca      `json:"tipoCobranca,omitempty"`
	SeuNumero             string            `json:"seuNumero,omitempty"`
	PessoaPagadora        string            `json:"pessoaPagadora,omitempty"`
	CpfCnpjPessoaPagadora string            `json:"cpfCnpjPessoaPagadora,omitempty"`
}

// SumarioItem representa um item de sumário de cobranças
type SumarioItem struct {
	Situacao   SituacaoCobranca `json:"situacao"`
	Valor      int64            `json:"valor"`
	Quantidade int64            `json:"quantidade"`
}

// Sumario busca o sumário de cobranças
func (s *Service) Sumario(ctx context.Context, request *SumarioRequest) (*[]SumarioItem, error) {
	token, err := s.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := s.backend.Req().
		SetContext(ctx).
		SetResult(&[]SumarioItem{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if request != nil {
		req.SetQueryParams(interutils.StructToMap(request))
	}

	resp, err := req.Post(path.Join(cobrancaEndpoint, "sumario"))
	if err != nil {
		return nil, erros.NewErrorWithStatus(resp.StatusCode(), resp.String())
	}

	// Check for errors
	if resp.IsError() {
		errResp, ok := resp.Error().(*erros.Response)
		if ok {
			return nil, errResp.WithStatus(resp.StatusCode())
		}
		return nil, erros.NewErrorWithStatus(resp.StatusCode(), resp.String())
	}

	return resp.Result().(*[]SumarioItem), nil
}
