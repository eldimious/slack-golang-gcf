package messages

// BookRepository provides an abstraction on top of the book data source
type MessageRepository interface {
	SendMessage(*Message) (*Message, error)
}
