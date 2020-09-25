package ok_slice_uint64

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Uint64, violationReceiver bad.Emitter) ok_action.Enum
}
