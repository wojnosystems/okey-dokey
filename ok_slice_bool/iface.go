package ok_slice_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Bool, violationReceiver bad.MessageReceiver) ok_action.Enum
}
