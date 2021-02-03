package ok_bool

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultFalseFormat(definition *False, value optional.Bool) string {
	return "must be false"
}

type False struct {
	Format func(definition *False, value optional.Bool) string
}

func (m *False) Validate(value optional.Bool, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultFalseFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual bool) {
		if actual {
			violationReceiver.Emit(formatter(m, value))
		}
	})
	return ok_action.Continue
}
