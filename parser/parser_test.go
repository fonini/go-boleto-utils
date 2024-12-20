package parser

import (
	"github.com/fonini/go-boleto-utils/utils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// test that input matches the value we want. If not, report an error on t.
func testValue(t *testing.T, input string, want *utils.Boleto) {
	v, _ := Parse(input)

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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
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
				CodeType:          "DIGITABLE_LINE",
			},
		},
		{"73990.00004 00001.223320 90126.130344 4 00000000000000",
			&utils.Boleto{IssuerBankCode: "739",
				IssuerBankName:    "Banco Cetelem S.A.",
				Currency:          9,
				IssuerReserved1:   "00000",
				CheckDigit1:       4,
				IssuerReserved2:   "0000122332",
				CheckDigit2:       0,
				IssuerReserved3:   "9012613034",
				CheckDigit3:       4,
				GeneralCheckDigit: 4,
				DueDate:           time.Date(1997, 10, 7, 0, 0, 0, 0, loc),
				Amount:            0,
				CodeType:          "DIGITABLE_LINE",
			},
		},
		{"34191990600000005001092664672997197273480000",
			&utils.Boleto{IssuerBankCode: "341",
				IssuerBankName:    "Itaú Unibanco S.A.",
				Currency:          9,
				IssuerReserved1:   "10926",
				CheckDigit1:       3,
				IssuerReserved2:   "6467299719",
				CheckDigit2:       0,
				IssuerReserved3:   "7273480000",
				CheckDigit3:       5,
				GeneralCheckDigit: 1,
				DueDate:           time.Date(2024, 11, 20, 0, 0, 0, 0, loc),
				Amount:            5,
				CodeType:          "BARCODE",
			},
		},
		{"73994000000000000000000000001223329012613034",
			&utils.Boleto{IssuerBankCode: "739",
				IssuerBankName:    "Banco Cetelem S.A.",
				Currency:          9,
				IssuerReserved1:   "00000",
				CheckDigit1:       4,
				IssuerReserved2:   "0000122332",
				CheckDigit2:       0,
				IssuerReserved3:   "9012613034",
				CheckDigit3:       4,
				GeneralCheckDigit: 4,
				DueDate:           time.Date(1997, 10, 7, 0, 0, 0, 0, loc),
				Amount:            0,
				CodeType:          "BARCODE",
			},
		},
		{"74898992100000845361121577703702280000282105",
			&utils.Boleto{IssuerBankCode: "748",
				IssuerBankName:    "Banco Cooperativo Sicredi S.A.",
				Currency:          9,
				IssuerReserved1:   "11215",
				CheckDigit1:       6,
				IssuerReserved2:   "7770370228",
				CheckDigit2:       0,
				IssuerReserved3:   "0000282105",
				CheckDigit3:       6,
				GeneralCheckDigit: 8,
				DueDate:           time.Date(2024, 12, 5, 0, 0, 0, 0, loc),
				Amount:            845.36,
				CodeType:          "BARCODE",
			},
		},
		{input: "123456789",
			want: nil,
		},
		{input: "42297 03006 00002 695286 04014 412722 8 76500000003720",
			want: &utils.Boleto{IssuerBankCode: "422",
				IssuerBankName:    "Banco Safra S.A.",
				Currency:          9,
				IssuerReserved1:   "70300",
				CheckDigit1:       6,
				IssuerReserved2:   "0000269528",
				CheckDigit2:       6,
				IssuerReserved3:   "0401441272",
				CheckDigit3:       2,
				GeneralCheckDigit: 8,
				DueDate:           time.Date(2018, 9, 17, 0, 0, 0, 0, loc),
				Amount:            37.2,
				CodeType:          "DIGITABLE_LINE",
			},
		},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_GetBoletoType(t *testing.T) {
	tests := []struct {
		input string
		want  utils.BoletoType
	}{
		{"826700000035 645607980002 010002351038 822024116714",
			utils.Sanitation,
		},
		{
			"836800000033 380600863225 535337514090 100168807509",
			utils.ElectricityAndGas,
		},
		{
			"85860000000 4 83740385242 0 43070124241 5 85141630306 0",
			utils.GovernmentAgencies,
		},
		{
			"856500000026 056505152027 411292024030 335182000000",
			utils.GovernmentAgencies,
		},
		{
			"846800000008 550000791008 011193989719 924101544345",
			utils.Telecommunications,
		},
		{
			input: "73990.00004 00001.223320 90126.130344400000000000000",
			want:  utils.CreditCard,
		},
	}

	for _, tt := range tests {
		msg := GetBoletoType(tt.input)

		if tt.want != msg {
			t.Fatalf(`GetBoletoType("%s") = %q, want match for %#q`, tt.input, msg, tt.want)
		}
	}
}
