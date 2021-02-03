package ok_uint64

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired) string
}

func (m *IsRequired) Validate(value optional.Uint64, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfUnset(func() {
		violationReceiver.Emit(formatter(m))
	})
	return ok_action.Continue
}
