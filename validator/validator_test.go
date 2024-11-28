package validator

import (
	"testing"
)

// test that input matches the value we want. If not, report an error on t.
func testValue(t *testing.T, input string, want bool) {
	v := Validate(input)

	if v != want {
		t.Errorf("Validate(%v) failed", input)
	}
}

func TestValues_Validate(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"34191.75124 34567.871230 41234.560005 8 92850000026035",
			true,
		},
		{"23793.38128 60005.963347 21000.063301 1 74640000116037",
			true,
		},
		{"74891.11611 00172.302267 05522.671006 3 69050000017500",
			false,
		},
		{"00190000090333717600600639372176398960000008000",
			true,
		},
		{"46191110000000000002635057041010498940000096000",
			true,
		},
		{"48190.00003 00005.150396 31049.960144 9 98650000025736",
			true,
		},
		{"75691303670103467211159238450015997710000096210",
			true,
		},
		{"75691313670103467211159238450015997710000096210",
			false,
		},
		{"73990.00004 00001.223320 90126.130344 4 00000000000000",
			true,
		},
		{"73990.00005 00001.223320 90126.130344 4 00000000000000",
			false,
		},
		{"73990.00004 00001.223320 90126.130344 4 00000000000000",
			true,
		},
		{"73994000000000000000000000001223329012613034",
			true,
		},
		{"34191990600000005001092664672997197273480000",
			true,
		},
		{
			input: "2325435435",
			want:  false,
		},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}
