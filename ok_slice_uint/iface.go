package ok_slice_uint

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Uint, violationReceiver bad.Emitter) ok_action.Enum
}
