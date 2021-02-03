package ok_slice_uint8

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

type Definer interface {
	Validate(value []optional.Uint8, violationReceiver bad.Emitter) ok_action.Enum
}
