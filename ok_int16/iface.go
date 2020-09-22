package ok_int16

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value optional.Int16, violationReceiver bad.MessageReceiver) ok_action.Enum
}
