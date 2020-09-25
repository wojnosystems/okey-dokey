package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func Validate(value optional.Bool, definitions []Definer, violations bad.Emitter) {
	for _, definition := range definitions {
		action := definition.Validate(value, violations)
		if action != ok_action.Continue {
			return
		}
	}
}
