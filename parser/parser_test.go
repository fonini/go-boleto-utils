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
		t.Errorf("Parse(%v) returned an error: %v", input, err)
	}

	if diff := cmp.Diff(want, v); diff != "" {
		t.Errorf("Parse(%v) mismatch:\n%s", input, diff)
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
				IssuerBankName:    "Itaú Unibanco S.A.",
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
				IssuerBankName:    "Banco Cooperativo Sicredi S.A.",
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
		{"00190000090333717600600639372176398960000008000",
			&utils.Boleto{IssuerBankCode: "001",
				IssuerBankName:    "Banco do Brasil S.A.",
				Currency:          9,
				IssuerReserved1:   "00000",
				CheckDigit1:       9,
				IssuerReserved2:   "0333717600",
				CheckDigit2:       6,
				IssuerReserved3:   "0063937217",
				CheckDigit3:       6,
				GeneralCheckDigit: 3,
				DueDate:           time.Date(2024, 11, 10, 0, 0, 0, 0, loc),
				Amount:            80,
			},
		},
		{"46191110000000000002635057041010498940000096000",
			&utils.Boleto{IssuerBankCode: "461",
				IssuerBankName:    "Asaas",
				Currency:          9,
				IssuerReserved1:   "11100",
				CheckDigit1:       0,
				IssuerReserved2:   "0000000002",
				CheckDigit2:       6,
				IssuerReserved3:   "3505704101",
				CheckDigit3:       0,
				GeneralCheckDigit: 4,
				DueDate:           time.Date(2024, 11, 8, 0, 0, 0, 0, loc),
				Amount:            960,
			},
		},
		{"48190.00003 00005.150396 31049.960144 9 98650000025736",
			&utils.Boleto{IssuerBankCode: "481",
				IssuerBankName:    "Superlógica Sociedade de Crédito Direto S.A.",
				Currency:          9,
				IssuerReserved1:   "00000",
				CheckDigit1:       3,
				IssuerReserved2:   "0000515039",
				CheckDigit2:       6,
				IssuerReserved3:   "3104996014",
				CheckDigit3:       4,
				GeneralCheckDigit: 9,
				DueDate:           time.Date(2024, 10, 10, 0, 0, 0, 0, loc),
				Amount:            257.36,
			},
		},
		{"75691303670103467211159238450015997710000096210",
			&utils.Boleto{IssuerBankCode: "756",
				IssuerBankName:    "Banco Cooperativo do Brasil S.A. - Bancoob",
				Currency:          9,
				IssuerReserved1:   "13036",
				CheckDigit1:       7,
				IssuerReserved2:   "0103467211",
				CheckDigit2:       1,
				IssuerReserved3:   "5923845001",
				CheckDigit3:       5,
				GeneralCheckDigit: 9,
				DueDate:           time.Date(2024, 7, 8, 0, 0, 0, 0, loc),
				Amount:            962.10,
			},
		},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}
