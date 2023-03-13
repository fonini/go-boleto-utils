package parser

import (
	"github.com/fonini/go-boleto-utils/utils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// test that input matches the value we want. If not, report an error on t.
func testValue(t *testing.T, input string, want *utils.Boleto) {
	v, err := Parse(input)

	if err != nil {
		t.Errorf("Pix(%v) returned an error: %v", input, err)
	}

	if diff := cmp.Diff(want, v); diff != "" {
		t.Errorf("Pix(%v) mismatch:\n%s", input, diff)
	}
}

func TestValues_Parse(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")

	tests := []struct {
		input string
		want  *utils.Boleto
	}{
		{"34191.75124 34567.871230 41234.560005 8 92850000026035",
			&utils.Boleto{IssuerBankCode: "341",
				IssuerBankName:    "ITAÚ UNIBANCO S.A.",
				Currency:          9,
				IssuerReserved1:   "17512",
				CheckDigit1:       4,
				IssuerReserved2:   "3456787123",
				CheckDigit2:       0,
				IssuerReserved3:   "4123456000",
				CheckDigit3:       5,
				GeneralCheckDigit: 8,
				DueDate:           time.Date(2023, 3, 10, 0, 0, 0, 0, loc),
				Amount:            260.35,
			},
		},
		{"23793.38128 60005.963347 21000.063301 1 74640000116037",
			&utils.Boleto{IssuerBankCode: "237",
				IssuerBankName:    "Banco Bradesco S.A.",
				Currency:          9,
				IssuerReserved1:   "33812",
				CheckDigit1:       8,
				IssuerReserved2:   "6000596334",
				CheckDigit2:       7,
				IssuerReserved3:   "2100006330",
				CheckDigit3:       1,
				GeneralCheckDigit: 1,
				DueDate:           time.Date(2018, 3, 15, 0, 0, 0, 0, loc),
				Amount:            1160.37,
			},
		},
		{"74891.11612 00172.302267 05522.671006 3 69050000017500",
			&utils.Boleto{IssuerBankCode: "748",
				IssuerBankName:    "BANCO COOPERATIVO SICREDI S.A.",
				Currency:          9,
				IssuerReserved1:   "11161",
				CheckDigit1:       2,
				IssuerReserved2:   "0017230226",
				CheckDigit2:       7,
				IssuerReserved3:   "0552267100",
				CheckDigit3:       6,
				GeneralCheckDigit: 3,
				DueDate:           time.Date(2016, 9, 2, 0, 0, 0, 0, loc),
				Amount:            175,
			},
		},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}
