package banking

import "github.com/raniellyferreira/interbank-go/backend"

const (
	endpointBanking = "banking/v2"
)

type Service struct {
	backend *backend.BackendImplement
}

// NewService creates a new banking service
func NewService(client *backend.BackendImplement) *Service {
	return &Service{
		backend: client,
	}
}
