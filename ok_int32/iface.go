package ok_int32

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value optional.Int32, violationReceiver bad.MessageReceiver) ok_action.Enum
}
