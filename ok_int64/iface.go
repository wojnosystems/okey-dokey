package ok_int64

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

type Definer interface {
	Validate(value optional.Int64, violationReceiver bad.Emitter) ok_action.Enum
}
