package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

type Definer interface {
	Validate(value optional.Bool, violationReceiver bad.Emitter) ok_action.Enum
}
