package main

import (
	"fmt"
	"github.com/fonini/go-boleto-utils/validator"
)

func main() {
	fmt.Println(validator.Validate("42297.03006 00002.695286 56809.062427 1 67950000039229"))
}
