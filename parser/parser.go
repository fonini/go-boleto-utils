package parser

import (
	"errors"
	"github.com/fonini/go-boleto-utils/utils"
	"strconv"
	"time"
)

// Parse parses a digitable line into a Boleto struct
func Parse(digitableLine string) (*utils.Boleto, error) {
	line := utils.OnlyNumbers(digitableLine)

	if len(line) != 47 {
		return nil, errors.New("the typeable line must be 47 characters long")
	}

	var boleto utils.Boleto

	boleto.IssuerBankCode = utils.Substr(line, 0, 3)
	boleto.IssuerBankName = utils.Banks[boleto.IssuerBankCode]

	boleto.Currency, _ = strconv.Atoi(utils.Substr(line, 3, 1))

	boleto.IssuerReserved1 = utils.Substr(line, 4, 5)
	boleto.CheckDigit1, _ = strconv.Atoi(utils.Substr(line, 9, 1))

	boleto.IssuerReserved2 = utils.Substr(line, 10, 10)
	boleto.CheckDigit2, _ = strconv.Atoi(utils.Substr(line, 20, 1))

	boleto.IssuerReserved3 = utils.Substr(line, 21, 10)
	boleto.CheckDigit3, _ = strconv.Atoi(utils.Substr(line, 31, 1))

	boleto.GeneralCheckDigit, _ = strconv.Atoi(utils.Substr(line, 32, 1))

	dueDate, _ := strconv.Atoi(utils.Substr(line, 33, 4))
	dateFactor, err := time.Parse("2006-01-02 15:04:05", utils.BaseDate)
	if err != nil {
		return nil, err
	}
	boleto.DueDate = dateFactor.AddDate(0, 0, dueDate)

	amount, err := strconv.ParseFloat(utils.Substr(line, 37, 10), 64)
	if err != nil {
		return nil, err
	}
	boleto.Amount = amount / 100

	return &boleto, nil
}
