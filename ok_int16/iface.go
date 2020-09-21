package ok_int16

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *int16, violationReceiver bad.MessageReceiver) ok_action.Enum
}
