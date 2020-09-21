package ok_int

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *int, violationReceiver bad.MessageReceiver) ok_action.Enum
}
