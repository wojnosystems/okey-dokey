package ok_bool

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *bool, violationReceiver bad.MessageReceiver) ok_action.Enum
}
