package main

import (
	"fmt"
	"math"
)

// func main() {
// 	// the goal of the app is to get input from the user then create the future value of an investment.
// 	const inflationRate = 3.0
// 	investmentAmount, years, expectedReturn := 54350.0, 10.0, 7.0
// 	// using := because the value type is to be inferred by go
// 	// fv
// 	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
// 	frv := futureValue / math.Pow(1+inflationRate/100, 7)
// 	// output the value
// 	fmt.Println(futureValue)
// 	fmt.Println(frv)
// }

func main() {
	const inflationRate = 28.3
	var investmentAmount, years, expectedReturn float64

	// collect user input
	fmt.Println("Enter investment amount, number of years, expected return (separated by spaces):")
	_, err := fmt.Scan(&investmentAmount, &years, &expectedReturn)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	// calculate the future values
	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	var realFutureValue = futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value considering expected return: %.2f\n", futureValue)
	fmt.Printf("Real Future Value considering inflation rate: %.2f\n", realFutureValue)
}
