package main

import (
	"fmt"

	"github.com/alpakih/git-flow-example/calculation"
	"github.com/alpakih/git-flow-example/testing"
)

func main() {

	fmt.Println("git flow example")

	fmt.Println(fmt.Sprintf("calculate sum = %d", calculation.Sum(1, 1, 1)))

	fmt.Println(testing.PrintOut())
}
