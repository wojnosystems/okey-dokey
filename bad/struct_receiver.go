package bad

// Receivers receive a message that a field failed to validate
type StructReceiver interface {
	MemberReceiver(structName string) MemberReceiver
}
