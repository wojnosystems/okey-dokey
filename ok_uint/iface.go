package ok_uint

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *uint, violationReceiver bad.MessageReceiver) ok_action.Enum
}
