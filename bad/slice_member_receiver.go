package bad

// SliceMemberReceiver is a convenience implementation of MemberReceiver used in tests, but applications should make their own
// so that they can handle notifications about errors
type SliceMemberReceiver map[string][]string

// Receive accepts an error message
func (r *SliceMemberReceiver) MessageReceiver(memberId string) MessageReceiver {
	return &smrMessage{
		parent:   r,
		memberId: memberId,
	}
}

type smrMessage struct {
	parent   *SliceMemberReceiver
	memberId string
}

func (m *smrMessage) ReceiveMessage(message string) {
	(*m.parent)[m.memberId] = append((*m.parent)[m.memberId], message)
}
