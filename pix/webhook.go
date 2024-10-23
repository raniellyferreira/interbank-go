package pix

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// CriarWebhook cria um webhook para receber notificações de pix
func (c *Service) CriarWebhook(ctx context.Context, chave, webhookUrl string) (*EmptyResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&EmptyResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"webhookUrl": webhookUrl,
		})

	resp, err := req.Put(path.Join(pixEndpoint, "webhook", chave))
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

	return resp.Result().(*EmptyResponse), nil
}

// ConsultarWebhook consulta um webhook
func (c *Service) ConsultarWebhook(ctx context.Context, chave string) (*WebhookResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&WebhookResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(pixEndpoint, "webhook", chave))
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

	return resp.Result().(*WebhookResponse), nil
}

// DeletarWebhook deleta um webhook
func (c *Service) DeletarWebhook(ctx context.Context, chave string) error {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Delete(path.Join(pixEndpoint, "webhook", chave))
	if err != nil {
		return erros.NewErrorWithStatus(resp.StatusCode(), resp.String())
	}

	// Check for errors
	if resp.IsError() {
		errResp, ok := resp.Error().(*erros.Response)
		if ok {
			return errResp.WithStatus(resp.StatusCode())
		}
		return erros.NewErrorWithStatus(resp.StatusCode(), resp.String())
	}

	return nil
}

// ConsultarWebhookCallbacks consulta os eventos de um webhook
func (c *Service) ConsultarWebhookCallbacks(ctx context.Context, request *ConsultarWebhooksCallbacksRequest) (*CallbacksResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&CallbacksResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if request != nil {
		req.SetQueryParams(interutils.StructToMap(request))
	}

	resp, err := req.Get(path.Join(pixEndpoint, "webhook/callbacks"))
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

	return resp.Result().(*CallbacksResponse), nil
}
