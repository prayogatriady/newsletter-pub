package ping

import (
	models_ping "newsletter-pub/models/ping"
)

type PingService interface {
	Ping() (*models_ping.PingResponse, error)
}

type pingService struct {
	pingRepository PingRepository
}

// Dependency injection
func NewPingService(repo PingRepository) PingService {
	return &pingService{
		pingRepository: repo,
	}
}

func (s *pingService) Ping() (response *models_ping.PingResponse, err error) {

	response = &models_ping.PingResponse{
		Message: "PONG!!!",
	}

	return
}
