package ok_uint32

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *uint32, violationReceiver bad.MessageReceiver) ok_action.Enum
}
