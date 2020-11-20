package bad

// Emitter receives a message that a field failed to validate
type Emitter interface {
	// Emit is called when a validation error is encountered.
	// The message is the human-readable message about what the problem was and potentially what to do about it.
	Emit(message string)
}
