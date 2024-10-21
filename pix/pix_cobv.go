package pix

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// CriarCobrancaComVencimentoETxID - Cria uma cobrança imediata com vencimento e txID
func (c *Service) CriarCobrancaComVencimentoETxID(ctx context.Context, txID string, request *CobrancaComVencimentoRequest) (*CobrancaComVencimentoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaComVencimentoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Put(path.Join(pixEndpoint, "cobv", txID))
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

	return resp.Result().(*CobrancaComVencimentoResponse), nil
}

// ConsultarCobrancasComVencimento - Consulta cobranças imediatas com vencimento
func (c *Service) ConsultarCobrancasComVencimento(ctx context.Context, request *ConsultarCobrancasComVencimentoRequest) (*ConsultarCobrancasComVencimentoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&ConsultarCobrancasComVencimentoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if request != nil {
		req.SetQueryParams(interutils.StructToMap(request))
	}

	resp, err := req.Get(path.Join(pixEndpoint, "cobv"))
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

	return resp.Result().(*ConsultarCobrancasComVencimentoResponse), nil
}

// ConsultarCobrancaComVencimento - Consulta uma cobrança com vencimento
func (c *Service) ConsultarCobrancaComVencimento(ctx context.Context, txID string) (*CobrancaComVencimentoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaComVencimentoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(pixEndpoint, "cobv", txID))
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

	return resp.Result().(*CobrancaComVencimentoResponse), nil
}

// EditarCobrancaComVencimento - Edita uma cobrança com vencimento e txID
func (c *Service) EditarCobrancaComVencimento(ctx context.Context, txID string, request *CobrancaComVencimentoRequest) (*CobrancaComVencimentoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaComVencimentoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Patch(path.Join(pixEndpoint, "cobv", txID))
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

	return resp.Result().(*CobrancaComVencimentoResponse), nil
}
