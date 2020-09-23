package ok_slice

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Counter interface {
	Len() int
}

type Definer interface {
	Validate(value Counter, violationReceiver bad.MessageReceiver) ok_action.Enum
}
