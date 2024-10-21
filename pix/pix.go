package pix

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/backend"
	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

const pixEndpoint = "pix/v2"

type Service struct {
	backend *backend.BackendImplement
}

func NewService(client *backend.BackendImplement) *Service {
	return &Service{
		backend: client,
	}
}

// ConsultarDevolucao para consultar a devolução de um pix
func (c *Service) ConsultarDevolucao(ctx context.Context, endToEndId, uniqId string) (*DevolucaoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&DevolucaoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(pixEndpoint, "pix", endToEndId, "devolucao", uniqId))
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

	return resp.Result().(*DevolucaoResponse), nil
}

// SolicitarDevolucao para solicitar a devolução de um pix
func (c *Service) SolicitarDevolucao(ctx context.Context, request *SolicitarDevolucaoPixRequest) (*DevolucaoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&DevolucaoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Put(path.Join(pixEndpoint, "pix", request.EndToEndID, "devolucao", request.GetLocalUniqId()))
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

	return resp.Result().(*DevolucaoResponse), nil
}

// Consultar pix recebidos
func (c *Service) ConsultarRecebidos(ctx context.Context, request *RecebidosRequest) (*RecebidosResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&RecebidosResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetQueryParams(interutils.StructToMap(request))

	resp, err := req.Get(path.Join(pixEndpoint, "pix"))
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

	return resp.Result().(*RecebidosResponse), nil
}

// consultaPix para consultar um pix através de um determinado EndToEndId
func (c *Service) Consultar(ctx context.Context, endToEndId string) (*Response, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&Response{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(pixEndpoint, "pix", endToEndId))
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

	return resp.Result().(*Response), nil
}
