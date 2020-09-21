package examples

import (
	"github.com/stretchr/testify/assert"
	"okey-dokey/bad"
	"okey-dokey/ok_int"
	"okey-dokey/ok_string"
	"testing"
)

type model struct {
	Name *string
	Age  *int
}

type modelValidationDefs struct {
	Name ok_string.On
	Age  ok_int.On
}

var modelValidations = modelValidationDefs{
	Name: ok_string.On{
		Id: "name",
		Ensure: []ok_string.Definer{
			&ok_string.IsRequired{},
			&ok_string.LengthAtMost{
				Length: 10,
			},
		},
	},
	Age: ok_int.On{
		Id: "age",
		Ensure: []ok_int.Definer{
			&ok_int.IsRequired{},
			&ok_int.GreaterThanOrEqual{
				Value: 18,
			},
		},
	},
}

// This method should be provided by the library as a reflection implementation
// It just goes through all of the items and evaluates them, then finally, evaluates the entire struct
func (m model) Validate(receiver bad.MemberReceiver) {
	{
		mem := receiver.MessageReceiver(modelValidations.Name.Id)
		ok_string.Validate(m.Name, &modelValidations.Name, mem)
	}
	{
		mem := receiver.MessageReceiver(modelValidations.Age.Id)
		ok_int.Validate(m.Age, &modelValidations.Age, mem)
	}
}

func TestModel(t *testing.T) {
	cases := map[string]struct {
		m        model
		expected map[string][]string
	}{
		"ok": {
			m: model{
				Name: strAddr("chris"),
				Age:  intAddr(30),
			},
			expected: map[string][]string{},
		},
		"string missing": {
			m: model{
				Age: intAddr(30),
			},
			expected: map[string][]string{
				"name": []string{"is required"},
			},
		},
		"age missing": {
			m: model{
				Name: strAddr("chris"),
			},
			expected: map[string][]string{
				"age": []string{"is required"},
			},
		},
		"age too young": {
			m: model{
				Name: strAddr("chris"),
				Age:  intAddr(17),
			},
			expected: map[string][]string{
				"age": []string{"must be greater than or equal to 18"},
			},
		},
		"name too long": {
			m: model{
				Name: strAddr("chriswojno1"),
				Age:  intAddr(30),
			},
			expected: map[string][]string{
				"name": []string{"cannot have more than 10 characters"},
			},
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := bad.SliceMemberReceiver{}
			c.m.Validate(&actual)
			assert.Equal(t, bad.SliceMemberReceiver(c.expected), actual)
		})
	}
}

func strAddr(value string) *string {
	return &value
}

func intAddr(value int) *int {
	return &value
}
