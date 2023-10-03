package router

import (
	"newsletter-pub/api/v1/email"
	"newsletter-pub/api/v1/ping"
)

func InitEmail() email.EmailController {
	emailRepository := email.NewEmailRepository()
	emailService := email.NewEmailService(emailRepository)
	emailController := email.NewEmailController(emailService)
	return emailController
}

func InitPing() ping.PingController {
	pingRepository := ping.NewPingRepository()
	pingService := ping.NewPingService(pingRepository)
	pingController := ping.NewPingController(pingService)
	return pingController
}
