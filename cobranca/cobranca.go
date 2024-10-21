package cobranca

import (
	"context"

	"github.com/raniellyferreira/interbank-go/backend"
	"github.com/raniellyferreira/interbank-go/erros"
)

const cobrancaEndpoint = "cobranca/v3/cobrancas"

type Service struct {
	backend *backend.BackendImplement
}

func NewService(client *backend.BackendImplement) *Service {
	return &Service{
		backend: client,
	}
}

func (c *Service) Emitir(ctx context.Context, request *EmitirRequest) (*EmitirResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&EmitirResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Post(cobrancaEndpoint)
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

	return resp.Result().(*EmitirResponse), nil
}
