package parser

import (
	"errors"
	"fmt"
	"github.com/fonini/go-boleto-utils/utils"
	"strconv"
	"time"
)

const (
	DigitableLine  utils.BoletoCodeType = "DIGITABLE_LINE"
	Barcode        utils.BoletoCodeType = "BARCODE"
	Unknown        utils.BoletoCodeType = "UNKNOWN"
	BaseDateFormat                      = "2006-01-02 15:04:05"
)

// Parse parses a digitable line or a barcode into a Boleto struct
func Parse(code string) (*utils.Boleto, error) {
	line := utils.OnlyNumbers(code)

	codeType, err := GetCodeType(line)

	if err != nil {
		return nil, err
	}

	if codeType == Barcode {
		line = ConvertBarcodeToDigitableLine(line)
	}

	boleto, err := parseDigitableLine(line)

	if err != nil {
		return nil, err
	}

	boleto.CodeType = codeType

	return boleto, nil
}

func GetCodeType(code string) (utils.BoletoCodeType, error) {
	code = utils.OnlyNumbers(code)

	switch len(code) {
	case 44:
		return Barcode, nil
	case 46, 47, 48:
		return DigitableLine, nil
	default:
		return Unknown, errors.New("unknown code")
	}
}

func parseDigitableLine(line string) (*utils.Boleto, error) {
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

	dueDate, err := calculateDueDate(utils.Substr(line, 33, 4))
	if err != nil {
		return nil, err
	}
	boleto.DueDate = dueDate

	amount, err := parseAmount(utils.Substr(line, 37, 10))
	if err != nil {
		return nil, err
	}
	boleto.Amount = amount

	return &boleto, nil
}

func calculateDueDate(dueDateStr string) (time.Time, error) {
	dueDate, err := strconv.Atoi(dueDateStr)
	if err != nil {
		return time.Time{}, err
	}
	dateFactor, err := time.Parse(BaseDateFormat, utils.BaseDate)
	if err != nil {
		return time.Time{}, err
	}
	return dateFactor.AddDate(0, 0, dueDate), nil
}

func parseAmount(amountStr string) (float64, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, err
	}
	return amount / 100, nil
}

func ConvertBarcodeToDigitableLine(barcode string) string {
	block1 := barcode[0:4] + barcode[19:24]
	cd1 := utils.CalculateVerificationDigit(block1)

	block2 := barcode[24:34]
	cd2 := utils.CalculateVerificationDigit(block2)

	block3 := barcode[34:44]
	cd3 := utils.CalculateVerificationDigit(block3)

	return fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s%s%s",
		block1[:5], block1[5:], cd1,
		block2[:5], block2[5:], cd2,
		block3[:5], block3[5:], cd3,
		barcode[4:5], barcode[5:19],
	)
}
