package cobranca

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
	interutils "github.com/raniellyferreira/interbank-go/utils"
)

// CriarWebhook represents a response to create a webhook
func (s *Service) CriarWebhook(ctx context.Context, request *CriarWebhookRequest) error {
	token, err := s.backend.Token(ctx)
	if err != nil {
		return err
	}

	req := s.backend.Req().
		SetContext(ctx).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(request)

	resp, err := req.Put(path.Join(cobrancaEndpoint, "webhook"))
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

// ConsultarWebhook represents a response to get a webhook
func (s *Service) ConsultarWebhook(ctx context.Context) (*Webhook, error) {
	token, err := s.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := s.backend.Req().
		SetContext(ctx).
		SetResult(&Webhook{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Get(path.Join(cobrancaEndpoint, "webhook"))
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

	return resp.Result().(*Webhook), nil
}

// ConsultarWebhookCallbacks represents a response to get a webhook callbacks
func (s *Service) ConsultarWebhookCallbacks(ctx context.Context, request *ConsultarWebhookCallbacksRequest) (*WebhookCallbacksResponse, error) {
	token, err := s.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := s.backend.Req().
		SetContext(ctx).
		SetResult(&WebhookCallbacksResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if request != nil {
		req.SetQueryParams(interutils.StructToMap(request))
	}

	resp, err := req.Get(path.Join(cobrancaEndpoint, "webhook", "callbacks"))
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

// DeleteWebhook represents a response to delete a webhook
func (s *Service) DeleteWebhook(ctx context.Context) error {
	token, err := s.backend.Token(ctx)
	if err != nil {
		return err
	}

	req := s.backend.Req().
		SetContext(ctx).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	resp, err := req.Delete(path.Join(cobrancaEndpoint, "webhook"))
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
