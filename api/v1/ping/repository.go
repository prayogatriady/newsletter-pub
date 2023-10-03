package ping

type PingRepository interface {
}

type pingRepository struct {
}

// Dependency injection
func NewPingRepository() PingRepository {
	return &pingRepository{}
}
