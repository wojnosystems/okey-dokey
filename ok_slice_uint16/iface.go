package ok_slice_uint16

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Uint16, violationReceiver bad.Emitter) ok_action.Enum
}
