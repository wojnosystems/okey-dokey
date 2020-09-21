package bad

// SliceMessageReceiver is a convenience implementation of MessageReceiver used in tests, but applications should make their own
// so that they can handle notifications about errors
type SliceMessageReceiver []string

// Receive accepts an error message
func (r *SliceMessageReceiver) ReceiveMessage(message string) {
	*r = append(*r, message)
}
