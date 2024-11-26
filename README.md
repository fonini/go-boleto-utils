# go-boleto-utils

[![GoDoc](https://pkg.go.dev/badge/github.com/fonini/go-boleto-utils)](https://pkg.go.dev/github.com/fonini/go-boleto-utils)
[![Test Status](https://github.com/fonini/go-boleto-utils/workflows/tests/badge.svg)](https://github.com/fonini/go-boleto-utils/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/github/fonini/go-boleto-utils/graph/badge.svg?token=L8ZSJUCHFJ)](https://codecov.io/github/fonini/go-boleto-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/fonini/go-boleto-utils?force=true)](https://goreportcard.com/report/github.com/fonini/go-boleto-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

`go-boleto-utils` is a comprehensive Go library designed to simplify working with Brazilian bank slips (boletos). This utility package provides robust parsing and validation functionalities for digitable lines and barcodes, making it easier for developers to integrate boleto-related operations into their Go applications.

## Prerequisites

- Go 1.16 or higher
- Basic understanding of Brazilian banking document structures

## Installation

Install the library using Go modules with the following commands:

```sh
go get -u github.com/fonini/go-boleto-utils/parser
go get -u github.com/fonini/go-boleto-utils/validator
```

## Boleto Parser Usage

The parser allows you to extract comprehensive details from a boleto's digitable line or barcode:

### Using digitable line
```go
package main

import (
    "fmt"
    "github.com/fonini/go-boleto-utils/parser"
)

func main() {
    digitableLine := "34191.75124 34567.871230 41234.560005 8 92850000026035"
    result, err := parser.Parse(digitableLine)
    if err != nil {
        fmt.Println("Error parsing the digitable line:", err)
        return
    }
    
    fmt.Printf("Bank: %s (%s)\n", result.IssuerBankName, result.IssuerBankCode)
    fmt.Printf("Amount: R$ %.2f\n", result.Amount)
    fmt.Printf("Due Date: %s\n", result.DueDate.Format("2006-01-02"))
    fmt.Printf("Code Type: %s\n", result.CodeType) // DIGITABLE_LINE
}
```

### Using barcode

```go
package main

import (
    "fmt"
    "github.com/fonini/go-boleto-utils/parser"
)

func main() {
    barCode := "74898992100000845361121577703702280000282105"
    result, err := parser.Parse(barCode)
    if err != nil {
        fmt.Println("Error parsing the barcode:", err)
        return
    }

    fmt.Printf("Bank: %s (%s)\n", result.IssuerBankName, result.IssuerBankCode)
    fmt.Printf("Amount: R$ %.2f\n", result.Amount)
    fmt.Printf("Due Date: %s\n", result.DueDate.Format("2006-01-02"))
    fmt.Printf("Code Type: %s\n", result.CodeType) // BARCODE
}
```

### Parser Output Fields

- `IssuerBankCode`: Numeric code of the issuing bank
- `IssuerBankName`: Name of the issuing bank
- `Currency`: Monetary representation code
- `DueDate`: Expiration date of the bank slip
- `Amount`: Total amount of the bank slip
- `CodeType`: Type of the input code (DIGITABLE_LINE, BARCODE or UNKNOWN)

## Boleto Validator Usage

Quickly validate the integrity of a boleto's digitable line:

```go
package main

import (
	"fmt"
	"github.com/fonini/go-boleto-utils/validator"
)

func main() {
	digitableLine := "34191.75124 34567.871230 41234.560005 8 92850000026035"

	if validator.Validate(digitableLine) {
		fmt.Println("✅ The boleto is valid")
	} else {
		fmt.Println("❌ The boleto is not valid")
	}
}
```

## Testing

Run comprehensive tests using the following commands:

```sh
# Run all tests
go test ./...

# Run tests for a specific package
go test ./validator

# Run tests with verbose output
go test -v ./...
```

## Supported Banks

While the library aims to support multiple Brazilian banks, please check the documentation for the most up-to-date list of supported institutions.

## Limitations

- The library focuses on parsing and validation
- Does not handle boleto payment or generation
- Requires well-formed digitable lines

## License

This project is licensed under the MIT License. See the `LICENSE` file for complete details.

## Contact

For issues, questions, or contributions, please open an issue on the GitHub repository.