package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value optional.Bool, fieldName string, definitions []Definer, violations bad.MemberReceiver) {
	mem := violations.MessageReceiver(fieldName)
	for _, definition := range definitions {
		action := definition.Validate(value, mem)
		if action != ok_action.Continue {
			return
		}
	}
}
