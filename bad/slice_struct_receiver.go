package bad

// SliceMemberReceiver is a convenience implementation of MemberReceiver used in tests, but applications should make their own
// so that they can handle notifications about errors
type SliceStructReceiver map[string]SliceMemberReceiver

// Receive accepts an error message
func (r *SliceStructReceiver) MemberReceiver(structId string) MemberReceiver {
	nr := make(SliceMemberReceiver)
	(*r)[structId] = nr
	return &nr
}
