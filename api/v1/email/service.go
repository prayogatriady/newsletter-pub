package email

import (
	"newsletter-pub/drivers/googlepubsub"
	models_email "newsletter-pub/models/email"
)

type EmailService interface {
	SendEmail(data *models_email.EmailPayload) (*models_email.EmailResponse, error)
}

type emailService struct {
	emailRepository EmailRepository
}

// Dependency injection
func NewEmailService(repo EmailRepository) EmailService {
	return &emailService{
		emailRepository: repo,
	}
}

func (s *emailService) SendEmail(data *models_email.EmailPayload) (response *models_email.EmailResponse, err error) {

	// stringData, err := json.Marshal(data)
	// if err != nil {
	// 	return
	// }

	// messageId, err := googlepubsub.Publisher(string(stringData))
	// if err != nil {
	// 	return
	// }

	messageId, err := googlepubsub.PublisherLoop(data)
	if err != nil {
		return
	}

	response = &models_email.EmailResponse{
		MessageId: messageId,
	}

	return
}
