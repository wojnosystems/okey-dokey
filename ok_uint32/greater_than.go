package ok_uint32

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultGreaterThanFormat(definition *GreaterThan, value optional.Uint32) string {
	return fmt.Sprintf("must be greater than %d", definition.Value)
}

type GreaterThan struct {
	Format func(definition *GreaterThan, value optional.Uint32) string
	Value  uint32
}

func (m *GreaterThan) Validate(value optional.Uint32, violationReceiver bad.Emitter) ok_action.Enum {
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
