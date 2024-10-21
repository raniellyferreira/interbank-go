package banking

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// ExportarExtrato exports the account statement
func (c *Service) ExportarExtrato(ctx context.Context, dataInicio, dataFim string) (*ExportarExtratoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&ExportarExtratoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if dataInicio != "" {
		req.SetQueryParam("dataInicio", dataInicio)
	}

	if dataFim != "" {
		req.SetQueryParam("dataFim", dataFim)
	}

	resp, err := req.Get(path.Join(endpointBanking, "extrato", "exportar"))
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

	return resp.Result().(*ExportarExtratoResponse), nil
}

// ConsultarExtratoCompleto consults the account statement
func (c *Service) ConsultarExtratoCompleto(ctx context.Context, req *ConsultarExtratoCompletoRequest) (*ConsultarExtratoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	request := c.backend.Req().
		SetContext(ctx).
		SetResult(&ConsultarExtratoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	// Set query parameters
	if req != nil {
		request.SetQueryParams(interutils.StructToMap(req))
	}

	resp, err := request.Get(path.Join(endpointBanking, "extrato", "completo"))
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

	return resp.Result().(*ConsultarExtratoResponse), nil
}

// ConsultarExtrato consults the account statement
func (c *Service) ConsultarExtrato(ctx context.Context, dataInicio, dataFim string) (*ConsultarExtratoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&ConsultarExtratoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if dataInicio != "" {
		req.SetQueryParam("dataInicio", dataInicio)
	}

	if dataFim != "" {
		req.SetQueryParam("dataFim", dataFim)
	}

	resp, err := req.Get(path.Join(endpointBanking, "extrato"))
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

	return resp.Result().(*ConsultarExtratoResponse), nil
}
