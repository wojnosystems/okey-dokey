package ok_uint16

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value *uint16, definitions *On, violations bad.MessageReceiver) {
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, violations)
		if action != ok_action.Continue {
			return
		}
	}
}
