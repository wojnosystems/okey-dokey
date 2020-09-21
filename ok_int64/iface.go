package ok_int64

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *int64, violationReceiver bad.MessageReceiver) ok_action.Enum
}
