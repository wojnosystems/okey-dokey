package ok_uint16

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *uint16, violationReceiver bad.MessageReceiver) ok_action.Enum
}
