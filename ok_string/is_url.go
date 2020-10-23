package ok_string

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"net/url"
)

func defaultIsURL(definition *IsURL, value optional.String) string {
	return "is not a URL"
}

type IsURL struct {
	Format func(definition *IsURL, value optional.String) string
}

func (m *IsURL) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultIsURL
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	_, err := url.Parse(value.Value())
	if err != nil {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
