package ok_int64

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLessThanOrEqualFormat(definition *LessThanOrEqual, value optional.Int64) string {
	return fmt.Sprintf("must be less than or equal to %d", definition.Value)
}

type LessThanOrEqual struct {
	Format func(definition *LessThanOrEqual, value optional.Int64) string
	Value  int64
}

func (m *LessThanOrEqual) Validate(value optional.Int64, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultLessThanOrEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() > m.Value {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
