package ok_int64

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value *int64, definitions *On, violations bad.MessageReceiver) {
	for _, definition := range definitions.Ensure {
		action := definition.Validate(value, violations)
		if action != ok_action.Continue {
			return
		}
	}
}
