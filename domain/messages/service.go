package messages

// MessageService defines book service behavior.
type MessageService interface {
	SendMessage(*Message) error
	SendErrorMessage(*Message) error
	SendSuccessMessage(*Message) error
}

// Service struct handles message business logic tasks.
type Service struct {
	repository MessageService
}

func (svc *Service) SendMessage(message *Message) error {
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
