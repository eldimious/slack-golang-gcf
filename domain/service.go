package messages

// BookService defines book service behavior.
type MessageService interface {
	SendMessage(*Message) (*Message, error)
}

// Service struct handles book business logic tasks.
type Service struct {
	repository MessageService
}

func (svc *Service) SendMessage(message *Message) (*Message, error) {
	return svc.repository.SendMessage(message)
}

// NewService creates a new service struct
func NewService(repository MessageService) *Service {
	return &Service{repository: repository}
}
