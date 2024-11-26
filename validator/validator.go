package validator

import (
	"fmt"
	"github.com/fonini/go-boleto-utils/parser"
	"github.com/fonini/go-boleto-utils/utils"
)

func Validate(code string) bool {
	boleto, err := parser.Parse(code)

	if err != nil {
		return false
	}

	var blocks []string

	blocks = append(blocks, fmt.Sprintf("%s%d%s%d", boleto.IssuerBankCode, boleto.Currency, boleto.IssuerReserved1, boleto.CheckDigit1))
	blocks = append(blocks, fmt.Sprintf("%s%d", boleto.IssuerReserved2, boleto.CheckDigit2))
	blocks = append(blocks, fmt.Sprintf("%s%d", boleto.IssuerReserved3, boleto.CheckDigit3))

	return validateBlocks(blocks)
}

func validateBlocks(blocks []string) bool {
	var validCount = 0

	for _, block := range blocks {
		if utils.Mod10CheckDigit(block) {
			validCount++
		}
	}

	return validCount == len(blocks)
}
