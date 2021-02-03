package ok_bool

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultTrueFormat(definition *True, value optional.Bool) string {
	return "must be true"
}

type True struct {
	Format func(definition *True, value optional.Bool) string
}

func (m *True) Validate(value optional.Bool, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultTrueFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual bool) {
		if !actual {
			violationReceiver.Emit(formatter(m, value))
		}
	})
	return ok_action.Continue
}
