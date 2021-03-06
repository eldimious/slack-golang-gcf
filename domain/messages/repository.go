package messages

// MessageRepository provides an abstraction on top of the message data source
type MessageRepository interface {
	SendMessage(*Message) error
	SendSuccessMessage(*Message) error
	SendErrorMessage(*Message) error
}
