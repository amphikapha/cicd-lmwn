package main

import (
	"fmt"

	"github.com/amphikapha/cicd-lmwn/internal/calculator"
)

func main() {
	sum := calculator.Add(1, 2)
	fmt.Println("Hello world !!!", sum)
}
