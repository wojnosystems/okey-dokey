package ok_uint64

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *uint64, violationReceiver bad.MessageReceiver) ok_action.Enum
}
