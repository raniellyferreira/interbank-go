package banking

import (
	"context"
	"path"

	"github.com/raniellyferreira/interbank-go/erros"
)

// ConsultarSaldoResponse represents the response of the ConsultarSaldo method
type ConsultarSaldoResponse struct {
	Disponivel float64 `json:"disponivel"` // Saldo disponível
}

// ConsultarSaldo consults the balance of the account
func (c *Service) ConsultarSaldo(ctx context.Context, dataSaldo string) (*ConsultarSaldoResponse, error) {
	token, err := c.backend.Token(ctx)
	if err != nil {
		return nil, err
	}

	req := c.backend.Req().
		SetContext(ctx).
		SetResult(&ConsultarSaldoResponse{}).
		SetError(&erros.Response{}).
		SetAuthToken(token.GetAccessToken())

	if dataSaldo != "" {
		req.SetQueryParam("dataSaldo", dataSaldo)
	}

	resp, err := req.Get(path.Join(endpointBanking, "saldo"))
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

	return resp.Result().(*ConsultarSaldoResponse), nil
}
