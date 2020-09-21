package ok_string

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value *string, definitions *On, violations bad.MessageReceiver) {
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, violations)
		if action != ok_action.Continue {
			return
		}
	}
}
