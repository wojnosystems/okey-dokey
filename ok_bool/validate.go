package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value optional.Bool, definitions *On, violations bad.MemberReceiver) {
	mem := violations.MessageReceiver(definitions.Id)
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, mem)
		if action != ok_action.Continue {
			return
		}
	}
}
