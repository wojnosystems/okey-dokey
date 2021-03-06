package ok_bool

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value optional.Bool) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value optional.Bool) string
}

func (m *IsRequired) Validate(value optional.Bool, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
