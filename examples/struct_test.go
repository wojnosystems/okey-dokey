package examples

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/go-sorted-set"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_int"
	"github.com/wojnosystems/okey-dokey/ok_range"
	"github.com/wojnosystems/okey-dokey/ok_slice_string"
	"github.com/wojnosystems/okey-dokey/ok_string"
	"testing"
)

type dogModel struct {
	Name optional.String
	Age  optional.Int
}

type dogValidationDefs struct {
	Name ok_string.On
	Age  ok_int.On
}

// Validations defined separately so that they can be swapped out on the userModel, depending on the situation
// This pre-compiles the validations so that they exist only once. Then the validations, which have no state,
// are run on the userModel
var dogValidations = dogValidationDefs{
	Name: ok_string.On{
		Ensure: []ok_string.Definer{
			&ok_string.IsRequired{},
			&ok_string.LengthBetween{
				Between: ok_range.IntBetween(2, 20),
			},
		},
	},
	Age: ok_int.On{
		Ensure: []ok_int.Definer{
			&ok_int.IsRequired{},
			&ok_int.LessThan{
				Value: 25,
			},
		},
	},
}

// This method should be provided by the library as a reflection implementation
// It just goes through all of the items and evaluates them, then finally, evaluates the entire struct
func (v dogValidationDefs) Validate(on *dogModel, receiver bad.MemberEmitter) {
	ok_string.Validate(on.Name, &v.Name, receiver.Into("name"))
	ok_int.Validate(on.Age, &v.Age, receiver.Into("age"))
}

type userModel struct {
	Name            optional.String
	Age             optional.Int
	IceCreamFlavors []optional.String
	Pet             dogModel
}

type userValidationDefs struct {
	Name            ok_string.On
	Age             ok_int.On
	IceCreamFlavors ok_slice_string.On
}

// Validations defined separately so that they can be swapped out on the userModel, depending on the situation
// This pre-compiles the validations so that they exist only once. Then the validations, which have no state,
// are run on the userModel
var userValidations = userValidationDefs{
	Name: ok_string.On{
		Ensure: []ok_string.Definer{
			&ok_string.IsRequired{},
			&ok_string.LengthAtMost{
				Length: 10,
			},
		},
	},
	Age: ok_int.On{
		Ensure: []ok_int.Definer{
			&ok_int.IsRequired{},
			&ok_int.GreaterThanOrEqual{
				Value: 18,
			},
		},
	},
	IceCreamFlavors: ok_slice_string.On{
		Ensure: []ok_slice_string.Definer{
			&ok_slice_string.ItemCountBetween{
				Between: ok_range.IntBetween(2, 10),
			},
		},
		EnsureItems: []ok_string.Definer{
			&ok_string.OneOf{
				Only: sorted_set.NewString("raspberry", "vanilla", "chocolate", "pistachio").
					Sort(),
			},
		},
	},
}

// This method should be provided by the library as a reflection implementation
// It just goes through all of the items and evaluates them, then finally, evaluates the entire struct
func (v userValidationDefs) Validate(on *userModel, receiver bad.MemberEmitter) {
	ok_string.Validate(on.Name, &v.Name, receiver.Into("name"))
	ok_int.Validate(on.Age, &v.Age, receiver.Into("age"))
	ok_slice_string.Validate(on.IceCreamFlavors, &v.IceCreamFlavors, receiver.Into("iceCreamFlavors"))
	dogValidations.Validate(&on.Pet, receiver.Into("pet"))
}

func TestModel(t *testing.T) {
	cases := map[string]struct {
		m        userModel
		expected map[string][]string
	}{
		"ok": {
			m: userModel{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
		},
		"string missing": {
			m: userModel{
				Age: optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
			expected: map[string][]string{
				"user.name": {
					"is required",
				},
			},
		},
		"age missing": {
			m: userModel{
				Name: optional.StringFrom("chris"),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
			expected: map[string][]string{
				"user.age": {
					"is required",
				},
			},
		},
		"age too young": {
			m: userModel{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(17),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
			expected: map[string][]string{
				"user.age": {
					"must be greater than or equal to 18",
				},
			},
		},
		"name too long": {
			m: userModel{
				Name: optional.StringFrom("veryLongName"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
			expected: map[string][]string{
				"user.name": {
					"cannot have more than 10 characters",
				},
			},
		},
		"bad ice cream": {
			m: userModel{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("coffee"),
					optional.StringFrom("chocolate"),
					optional.StringFrom("orange"),
				},
				Pet: dogModel{
					Name: optional.StringFrom("Zoey"),
					Age:  optional.IntFrom(5),
				},
			},
			expected: map[string][]string{
				"user.iceCreamFlavors[0]": {
					"must be one of the following: chocolate, pistachio, raspberry, vanilla",
				},
				"user.iceCreamFlavors[2]": {
					"must be one of the following: chocolate, pistachio, raspberry, vanilla",
				},
			},
		},
		"missing pet": {
			m: userModel{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{
				"user.pet.name": {
					"is required",
				},
				"user.pet.age": {
					"is required",
				},
			},
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := bad.NewCollection()
			userValidations.Validate(&c.m, actual.Into("user"))
			if c.expected == nil || len(c.expected) == 0 {
				assert.True(t, actual.IsEmpty())
			} else {
				assert.True(t, actual.HasAny())
				assert.Equal(t, c.expected, collectorToMap(actual))
			}
		})
	}
}

func collectorToMap(c bad.Collector) (out map[string][]string) {
	out = make(map[string][]string)
	for _, path := range c.Paths() {
		out[path] = c.MessagesAtPath(path)
	}
	return
}
