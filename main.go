package main

import (
	"fmt"

	"github.com/alpakih/git-flow-example/calculation"
	"github.com/alpakih/git-flow-example/testing"
	"github.com/alpakih/git-flow-example/testing2"
)

func main() {

	fmt.Println("git flow example")

	fmt.Println(fmt.Sprintf("calculate sum = %d", calculation.Sum(1, 1, 1)))

	fmt.Println(testing.PrintOut())

	fmt.Println(testing.PrintOutTest())

	fmt.Println(testing2.PrintOutNew())
}
