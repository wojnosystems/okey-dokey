package ok_slice_int32

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Int32, violationReceiver bad.Emitter) ok_action.Enum
}
