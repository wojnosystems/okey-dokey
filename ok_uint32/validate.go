package ok_uint32

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value optional.Uint32, on *On, violations bad.MemberReceiver) {
	mem := violations.MessageReceiver(on.Id)
	for _, definition := range on.Ensure {
		action := definition.Validate(value, mem)
		if action != ok_action.Continue {
			return
		}
	}
}
