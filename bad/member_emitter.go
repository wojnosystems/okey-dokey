package bad

type MemberEmitter interface {
	Into(memberId string) MemberEmitter
	Emitter
}
