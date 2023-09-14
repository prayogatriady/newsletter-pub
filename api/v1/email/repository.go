package email

type EmailRepository interface {
}

type emailRepository struct {
}

// Dependency injection
func NewEmailRepository() EmailRepository {
	return &emailRepository{}
}
