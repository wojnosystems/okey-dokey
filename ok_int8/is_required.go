package ok_int8

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value optional.Int8) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value optional.Int8) string
}

func (m *IsRequired) Validate(value optional.Int8, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
