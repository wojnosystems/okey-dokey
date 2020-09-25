package ok_uint32

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultEqualFormat(definition *Equal, value optional.Uint32) string {
	return fmt.Sprintf("must be exactly %d", definition.Value)
}

type Equal struct {
	Format func(definition *Equal, value optional.Uint32) string
	Value  uint32
}

func (m *Equal) Validate(value optional.Uint32, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() != m.Value {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
