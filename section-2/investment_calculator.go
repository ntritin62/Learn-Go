package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:",futureValue)
	fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue,futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):",futureRealValue)
}