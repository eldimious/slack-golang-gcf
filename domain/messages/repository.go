package messages

// MessageRepository provides an abstraction on top of the message data source
type MessageRepository interface {
	SendMessage(*Message) (string, error)
	SendSuccessMessage(*Message) (string, error)
	SendErrorMessage(*Message) (string, error)
}
