package messages

// MessageDispatcher defines an interface for sending messages
type MessageDispatcher interface {
	SendMessage(message *Message) (*string, error)
}

// MessageService defines message service behavior.
type MessageService interface {
	MessageDispatcher
}

// Service struct handles messagess business logic tasks.
type Service struct {
	MessageDispatcher
}

// NewService creates a new service struct
func NewService(dispatcher MessageDispatcher) *Service {
	return &Service{MessageDispatcher: dispatcher}
}
