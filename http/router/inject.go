package router

import "newsletter-pub/api/v1/email"

func InitEmail() email.EmailController {
	emailRepository := email.NewEmailRepository()
	emailService := email.NewEmailService(emailRepository)
	emailController := email.NewEmailController(emailService)
	return emailController
}
