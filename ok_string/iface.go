package ok_string

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum
}
