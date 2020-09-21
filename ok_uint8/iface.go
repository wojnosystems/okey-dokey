package ok_uint8

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *uint8, violationReceiver bad.MessageReceiver) ok_action.Enum
}
