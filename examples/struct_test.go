package examples

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-optional"
	sorted_set "github.com/wojnosystems/go-sorted-set"
	"okey-dokey/bad"
	"okey-dokey/ok_int"
	"okey-dokey/ok_range"
	"okey-dokey/ok_slice_string"
	"okey-dokey/ok_string"
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
func (v dogValidationDefs) Validate(on *dogModel, receiver bad.MemberReceiver) {
	ok_string.Validate(on.Name, &v.Name, receiver.MemberReceiver("name"))
	ok_int.Validate(on.Age, &v.Age, receiver.MemberReceiver("age"))
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
func (v userValidationDefs) Validate(on *userModel, receiver bad.MemberReceiver) {
	ok_string.Validate(on.Name, &v.Name, receiver.MemberReceiver("name"))
	ok_int.Validate(on.Age, &v.Age, receiver.MemberReceiver("age"))
	ok_slice_string.Validate(on.IceCreamFlavors, &v.IceCreamFlavors, receiver.MemberReceiver("iceCreamFlavors"))
	dogValidations.Validate(&on.Pet, receiver.MemberReceiver("pet"))
}

func TestModel(t *testing.T) {
	cases := map[string]struct {
		m        userModel
		expected bad.MemberReceiver
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
			expected: bad.NewSliceMemberReceiver(),
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.name": {
						"is required",
					},
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.age": {
						"is required",
					},
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.age": {
						"must be greater than or equal to 18",
					},
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.name": {
						"cannot have more than 10 characters",
					},
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.iceCreamFlavors[0]": {
						"must be one of the following: chocolate, pistachio, raspberry, vanilla",
					},
					"user.iceCreamFlavors[2]": {
						"must be one of the following: chocolate, pistachio, raspberry, vanilla",
					},
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
			expected: &bad.SliceMember{
				BadFields: map[string][]string{
					"user.pet.name": {
						"is required",
					},
					"user.pet.age": {
						"is required",
					},
				},
			},
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := bad.NewSliceMemberReceiver()
			userValidations.Validate(&c.m, actual.MemberReceiver("user"))
			assert.Equal(t, c.expected, actual)
		})
	}
}
