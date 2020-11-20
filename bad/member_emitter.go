package bad

// MemberEmitter is just like Emitter, but tracks when the validation algorithm descends Into a field
type MemberEmitter interface {
	Into(fieldName string) MemberEmitter
	Emitter
}
