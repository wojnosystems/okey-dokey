package ok_uint64

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultGreaterThanFormat(definition *GreaterThan, value uint64) string {
	return fmt.Sprintf("must be greater than %d", definition.Value)
}

type GreaterThan struct {
	Format func(definition *GreaterThan, value uint64) string
	Value  uint64
}

func (m *GreaterThan) Validate(value optional.Uint64, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultGreaterThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual uint64) {
		if actual <= m.Value {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
