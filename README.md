# go-boleto-utils

[![GoDoc](https://pkg.go.dev/badge/github.com/fonini/go-boleto-utils)](https://pkg.go.dev/github.com/fonini/go-boleto-utils)
[![Test Status](https://github.com/fonini/go-boleto-utils/workflows/tests/badge.svg)](https://github.com/fonini/go-boleto-utils/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/github/fonini/go-boleto-utils/graph/badge.svg?token=L8ZSJUCHFJ)](https://codecov.io/github/fonini/go-boleto-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/fonini/go-boleto-utils?force=true)](https://goreportcard.com/report/github.com/fonini/go-boleto-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Objective

This project aims to provide functionality for validation and parsing of Brazilian bank slips (boletos).

## Installation

To install this project, you can use the following command:

```sh
go get -u github.com/fonini/go-boleto-utils/parser
go get -u github.com/fonini/go-boleto-utils/validator
```

## Usage of Boleto Parser

You can use the parser as follows:

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

    fmt.Println(result)

    // Output: {"IssuerBankCode":"341","IssuerBankName":"Itaú Unibanco S.A.","Currency":9,"IssuerReserved1":"17512","CheckDigit1":4,"IssuerReserved2":"3456787123","CheckDigit2":0,"IssuerReserved3":"4123456000","CheckDigit3":5,"GeneralCheckDigit":8,"DueDate":"2023-03-10T00:00:00Z","Amount":260.35}
}
```

## Usage of Boleto Validator

You can use the validator as follows:

```go
package main

import (
    "fmt"
    "github.com/fonini/go-boleto-utils/validator"
)

func main() {
    digitableLine := "34191.75124 34567.871230 41234.560005 8 92850000026035"
    isValid := validator.Validate(digitableLine)

    if isValid {
        fmt.Println("The boleto is valid")
    } else {
        fmt.Println("The boleto is not valid")
    }
}
```

## Running Tests

To run the tests for this project, you can use the following command:

```sh
go test ./...
```

This command will run all the tests in the project, including those for the parser and validator packages.

If you want to run a specific test file or package, you can specify the path:

```sh
go test ./validator
```

Or to run tests in a specific file:

```sh
go test ./validator/validator_test.go
```

Make sure to have Go installed and properly configured in your environment to run the tests correctly.

## License

This project is licensed under the terms of the MIT license. See the `LICENSE` file for more details.