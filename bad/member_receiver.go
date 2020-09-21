package bad

// Receivers receive a message that a field failed to validate
type MemberReceiver interface {
	MessageReceiver(memberId string) MessageReceiver
}
