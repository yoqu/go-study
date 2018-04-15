package main

import "fmt"

func main() {
	const freezingF, boilingF = 21.0, 212.0
	fmt.Printf("%g = %g\n", freezingF, ftoC(freezingF))
	fmt.Printf("%g =%g \n", boilingF, ftoC(boilingF))
}

func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
