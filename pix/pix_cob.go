package pix

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// EditarCobrancaImediata edita uma cobrança imediata.
func (c *Service) EditarCobrancaImediata(ctx context.Context, txID string, request *CobrancaImediataRequest) (*CobrancaImediataResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaImediataResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Patch(path.Join(pixEndpoint, "cob", txID))
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

	return resp.Result().(*CobrancaImediataResponse), nil
}

// ConsultarCobrancasImediatas consulta cobranças imediatas.
func (c *Service) ConsultarCobrancasImediatas(ctx context.Context, request *ConsultarCobrancasImediatasRequest) (*ConsultarCobrancasImediatasResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&ConsultarCobrancasImediatasResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if request != nil {
		req.SetQueryParams(interutils.StructToMap(request))
	}

	resp, err := req.Get(path.Join(pixEndpoint, "cob"))
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

	return resp.Result().(*ConsultarCobrancasImediatasResponse), nil
}

// ConsultarCobrancaImediata consulta uma cobrança imediata.
func (c *Service) ConsultarCobrancaImediata(ctx context.Context, txID string) (*CobrancaImediataResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaImediataResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(pixEndpoint, "cob", txID))
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

	return resp.Result().(*CobrancaImediataResponse), nil
}

// CriarCobrancaImediataComTxID cria uma cobrança imediata com o txID informado.
func (c *Service) CriarCobrancaImediataComTxID(ctx context.Context, txID string, request *CobrancaImediataRequest) (*CobrancaImediataResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaImediataResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Put(path.Join(pixEndpoint, "cob", txID))
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

	return resp.Result().(*CobrancaImediataResponse), nil
}

// CriarCobrancaImediata cria uma cobrança imediata.
func (c *Service) CriarCobrancaImediata(ctx context.Context, request *CobrancaImediataRequest) (*CobrancaImediataResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CobrancaImediataResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Post(path.Join(pixEndpoint, "cob"))
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

	return resp.Result().(*CobrancaImediataResponse), nil
}

// PagarCobrancaImediata paga uma cobrança imediata. (SandBox apenas)
func (c *Service) PagarCobrancaImediata(ctx context.Context, tipoCob TipoCobranca, txID, valor string) (*map[string]interface{}, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&map[string]interface{}{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"valor": valor,
		})

	resp, err := req.Post(path.Join(pixEndpoint, string(tipoCob), "pagar", txID))
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

	return resp.Result().(*map[string]interface{}), nil
}
