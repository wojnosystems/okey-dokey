package ok_int16

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultGreaterThanFormat(definition *GreaterThan, value optional.Int16) string {
	return fmt.Sprintf("must be greater than %d", definition.Value)
}

type GreaterThan struct {
	Format func(definition *GreaterThan, value optional.Int16) string
	Value  int16
}

func (m *GreaterThan) Validate(value optional.Int16, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultGreaterThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() <= m.Value {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
