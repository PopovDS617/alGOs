package search

import (
	"testing"
)

type testExampleStandard struct {
	Email     string
	Character struct {
		Nickname string
		Level    int
	}
}

type testExampleNested struct {
	HighA struct {
		MiddleA struct {
			LowA struct {
				Lowest struct {
					Color string
				}
			}
		}
	}
	HighB struct {
		MiddleB struct {
			LowB struct {
				Lowest struct {
					Color string
				}
			}
		}
	}
	HighC struct {
		MiddleC struct {
			LowC struct {
				Lowest struct {
					Color string
				}
			}
		}
	}
	HighD struct {
		MiddleD struct {
			LowD struct {
				Lowest struct {
					Color string
				}
			}
		}
	}
	HighE struct {
		MiddleE struct {
			LowE struct {
				Lowest struct {
					Color string
				}
			}
		}
	}
}

var exampleNested testExampleNested
var exampleStandard testExampleStandard

func TestMain(m *testing.M) {

	exampleStandard = testExampleStandard{
		Email: "hello@there.com",
		Character: struct {
			Nickname string
			Level    int
		}{
			Nickname: "spellsinger",
			Level:    70,
		},
	}

	exampleNested.HighA.MiddleA.LowA.Lowest.Color = "Red"
	exampleNested.HighB.MiddleB.LowB.Lowest.Color = "Black"
	exampleNested.HighC.MiddleC.LowC.Lowest.Color = "Yellow"
	exampleNested.HighD.MiddleD.LowD.Lowest.Color = "Orange"
	exampleNested.HighE.MiddleE.LowE.Lowest.Color = "Cyan"

	m.Run()
}

func TestUseBFS(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		obj      any
		expected bool
	}{
		{"nested_success", "Color", "Yellow", exampleNested, true},
		{"nested_failure", "Color", "Green", exampleNested, false},
		{"standard_success", "Color", "Yellow", exampleStandard, false},
		{"standard_failure", "Color", "Green", exampleStandard, false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			res := UseBFS(test.key, test.value, test.obj)
			if res != test.expected {
				t.Errorf("expected %v to be %v", res, test.expected)
			}
		})

	}

}
func TestUseDFS(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		obj      any
		expected bool
	}{
		{"nested_success", "Color", "Yellow", exampleNested, true},
		{"nested_failure", "Color", "Green", exampleNested, false},
		{"standard_success", "Color", "Yellow", exampleStandard, false},
		{"standard_failure", "Color", "Green", exampleStandard, false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			res := UseDFS(test.key, test.value, test.obj)
			if res != test.expected {
				t.Errorf("expected %v to be %v", res, test.expected)
			}
		})

	}

}
func TestUseDFSRecursively(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		obj      any
		expected bool
	}{
		{"nested_success", "Color", "Yellow", exampleNested, true},
		{"nested_failure", "Color", "Green", exampleNested, false},
		{"standard_success", "Color", "Yellow", exampleStandard, false},
		{"standard_failure", "Color", "Green", exampleStandard, false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			res := UseDFSRecursively(test.key, test.value, test.obj)
			if res != test.expected {
				t.Errorf("expected %v to be %v", res, test.expected)
			}
		})

	}

}

func BenchmarkUseBFS_Not_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseBFS("Color", "Yellow", exampleStandard)
	}

}
func BenchmarkUseBFS_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseBFS("Color", "Yellow", exampleNested)
	}

}
func BenchmarkUseBFS_Nested_Value_Not_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseBFS("Color", "Transparent", exampleNested)
	}

}

func BenchmarkUseDFS_Not_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFS("Color", "Yellow", exampleStandard)
	}
}
func BenchmarkUseDFS_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFS("Color", "Yellow", exampleNested)
	}
}

func BenchmarkUseDFS_Nested_Value_Not_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFS("Color", "Transparent", exampleNested)
	}

}

func BenchmarkUseDFSRecursively_Not_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFSRecursively("Color", "Yellow", exampleStandard)
	}

}
func BenchmarkUseDFSRecursively_Nested_Value_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFSRecursively("Color", "Yellow", exampleNested)
	}

}
func BenchmarkUseDFSRecursively_Nested_Value_Not_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseDFSRecursively("Color", "Transparent", exampleNested)
	}

}
