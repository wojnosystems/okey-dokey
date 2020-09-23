package ok_uint16

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value optional.Uint16, definitions *On, violations bad.MemberReceiver) {
	mem := violations.MessageReceiver(definitions.Id)
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, mem)
		if action != ok_action.Continue {
			return
		}
	}
}
