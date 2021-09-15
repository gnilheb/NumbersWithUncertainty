package main

import (
	"fmt"

	"github.com/gnilheb/numberswithuncertainty/pkg/number"
)

func main() {
	fmt.Println("This is working")
	centralValue := 16.0
	uncertainty := 2.1
	a := number.NumberWithUncertainty{
		Number:           centralValue,
		Uncertainty:      uncertainty,
		DistributionType: number.None,
	}
	a.Init()

	b, err := a.Value()
	if err != nil {
		fmt.Printf("Error from call to Value() : %s", err.Error())
	}

	fmt.Printf("sampled value : %f\n", b)
}
