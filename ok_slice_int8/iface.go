package ok_slice_int8

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Int8, violationReceiver bad.Emitter) ok_action.Enum
}
