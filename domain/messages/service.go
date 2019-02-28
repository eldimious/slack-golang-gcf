package messages

// BookService defines book service behavior.
type MessageService interface {
	SendMessage(*Message) (string, error)
	SendErrorMessage(*Message) (string, error)
	SendSuccessMessage(*Message) (string, error)
}

// Service struct handles book business logic tasks.
type Service struct {
	repository MessageService
}

func (svc *Service) SendMessage(message *Message) (string, error) {
	if message.Type == "error" {
		return svc.repository.SendErrorMessage(message)
	} else if message.Type == "success" {
		return svc.repository.SendSuccessMessage(message)
	} else {
		return svc.repository.SendMessage(message)
	}
}

// NewService creates a new service struct
func NewService(repository MessageRepository) *Service {
	return &Service{repository: repository}
}
