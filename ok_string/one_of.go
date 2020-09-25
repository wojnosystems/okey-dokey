package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	sorted_set "github.com/wojnosystems/go-sorted-set"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"sort"
	"strings"
)

const (
	oneOfMaxItemsToList = 10
)

func defaultFormatOneOf(definition *OneOf, value optional.String) string {
	itemsToList := oneOfMaxItemsToList
	if len(definition.Only) < itemsToList {
		itemsToList = len(definition.Only)
	}
	list := strings.Join(definition.Only[0:itemsToList], ", ")
	if itemsToList < len(definition.Only) {
		list += ", ..."
	}
	return fmt.Sprintf("must be one of the following: %s", list)
}

type OneOf struct {
	Format func(definition *OneOf, value optional.String) string
	Only   sorted_set.String
}

func (m *OneOf) Validate(value optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultFormatOneOf
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	searchIndex := sort.SearchStrings(m.Only, value.Value())
	if searchIndex == len(m.Only) || m.Only[searchIndex] != value.Value() {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
