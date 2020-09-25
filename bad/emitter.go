package bad

// Receivers receive a message that a field failed to validate
type Emitter interface {
	Emit(message string)
}
