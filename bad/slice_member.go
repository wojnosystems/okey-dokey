package bad

import "strings"

// SliceMemberReceiver is a convenience implementation of MemberReceiver used in tests, but applications should make their own
// so that they can handle notifications about errors
type Fields struct {
	parent      *Fields
	currentName string
	BadFields   map[string][]string
}

func NewFields() *Fields {
	return &Fields{
		BadFields: make(map[string][]string),
	}
}

func (r *Fields) down(memberId string) *Fields {
	name := memberId
	if r.currentName != "" {
		name = r.currentName
		if !strings.HasPrefix(memberId, "[") {
			name += "."
		}
		name += memberId
	}
	return &Fields{
		currentName: name,
		parent:      r,
		BadFields:   r.BadFields,
	}
}

func (r *Fields) MemberReceiver(memberId string) MemberReceiver {
	return r.down(memberId)
}

// Receive accepts an error message
func (r *Fields) MessageReceiver(memberId string) MessageReceiver {
	return r.down(memberId)
}

func (r *Fields) ReceiveMessage(message string) {
	r.BadFields[r.currentName] = append(r.BadFields[r.currentName], message)
}
