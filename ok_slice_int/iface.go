package ok_slice_int

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Int, violationReceiver bad.Emitter) ok_action.Enum
}
