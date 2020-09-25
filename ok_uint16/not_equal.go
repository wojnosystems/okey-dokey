package ok_uint16

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultNotEqualFormat(definition *NotEqual, value optional.Uint16) string {
	return fmt.Sprintf("must not be %d", definition.Value)
}

type NotEqual struct {
	Format func(definition *NotEqual, value optional.Uint16) string
	Value  uint16
}

func (m *NotEqual) Validate(value optional.Uint16, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultNotEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() == m.Value {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
