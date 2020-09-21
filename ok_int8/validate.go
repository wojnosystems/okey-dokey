package ok_int8

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value *int8, definitions *On, violations bad.MessageReceiver) {
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, violations)
		if action != ok_action.Continue {
			return
		}
	}
}
