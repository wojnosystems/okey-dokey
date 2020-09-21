package ok_int8

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *int8, violationReceiver bad.MessageReceiver) ok_action.Enum
}
