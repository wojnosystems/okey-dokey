package ok_uint32

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

type Definer interface {
	Validate(value optional.Uint32, violationReceiver bad.Emitter) ok_action.Enum
}
