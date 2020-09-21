package ok_int32

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *int32, violationReceiver bad.MessageReceiver) ok_action.Enum
}
