package bad

// Receivers receive a message that a field failed to validate
type MessageReceiver interface {
	ReceiveMessage(message string)
}
