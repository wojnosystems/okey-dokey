package bad

import "strings"

// SliceMemberReceiver is a convenience implementation of MemberReceiver used in tests, but applications should make their own
// so that they can handle notifications about errors
type SliceMember struct {
	parent      *SliceMember
	currentName string
	BadFields   map[string][]string
}

func NewSliceMemberReceiver() *SliceMember {
	return &SliceMember{
		BadFields: make(map[string][]string),
	}
}

func (r *SliceMember) down(memberId string) *SliceMember {
	name := memberId
	if r.currentName != "" {
		name = r.currentName
		if !strings.HasPrefix(memberId, "[") {
			name += "."
		}
		name += memberId
	}
	return &SliceMember{
		currentName: name,
		parent:      r,
		BadFields:   r.BadFields,
	}
}

func (r *SliceMember) MemberReceiver(memberId string) MemberReceiver {
	return r.down(memberId)
}

// Receive accepts an error message
func (r *SliceMember) MessageReceiver(memberId string) MessageReceiver {
	return r.down(memberId)
}

func (r *SliceMember) ReceiveMessage(message string) {
	r.BadFields[r.currentName] = append(r.BadFields[r.currentName], message)
}
