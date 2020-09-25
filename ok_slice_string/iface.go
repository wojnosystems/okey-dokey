package ok_slice_string

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.String, violationReceiver bad.Emitter) ok_action.Enum
}
