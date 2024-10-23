package banking

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// CriarWebhook cria um webhook para receber notificações de pix ou boleto
func (c *Service) CriarWebhook(ctx context.Context, tipo TipoWebhook, webhookUrl string) error {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"webhookUrl": webhookUrl,
		})

	resp, err := req.Put(path.Join(endpointBanking, "webhooks", string(tipo)))
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

// ConsultarWebhook consulta um webhook
func (c *Service) ConsultarWebhook(ctx context.Context, tipo TipoWebhook) (*WebhookResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&WebhookResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(endpointBanking, "webhooks", string(tipo)))
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
func (c *Service) DeletarWebhook(ctx context.Context, tipo TipoWebhook) error {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Delete(path.Join(endpointBanking, "webhooks", string(tipo)))
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

// ConsultarWebhooksCallbacks consulta os eventos de um webhook
func (c *Service) ConsultarWebhooksCallbacks(ctx context.Context, tipo TipoWebhook, req *WebhookCallbacksRequest) (*WebhookCallbacksResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	reqConsulta := c.backend.Req().
		SetContext(ctx).
		SetResult(&WebhookCallbacksResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if req != nil {
		reqConsulta.SetQueryParams(interutils.StructToMap(req))
	}

	resp, err := reqConsulta.Get(path.Join(endpointBanking, "webhooks", string(tipo), "callbacks"))
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

	return resp.Result().(*WebhookCallbacksResponse), nil
}
