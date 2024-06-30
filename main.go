package main

import "fmt"

const (
	BubblegumTitle     = "Bubblegum"
	BubblegumTotal     = 202
	ToffeeTitle        = "Toffee"
	ToffeeTotal        = 118
	IceCreamTitle      = "Ice cream"
	IceCreamTotal      = 2250
	MilkChocolateTitle = "Milk chocolate"
	MilkChocolateTotal = 1680
	DoughnutTitle      = "Doughnut"
	DoughnutTotal      = 1075
	PancakeTitle       = "Pancake"
	PancakeTotal       = 80
)

func main() {
	// Write your code here
	fmt.Println("Earned amount:")
	fmt.Printf("%s: $%d\n", BubblegumTitle, BubblegumTotal)
	fmt.Printf("%s: $%d\n", ToffeeTitle, ToffeeTotal)
	fmt.Printf("%s: $%d\n", IceCreamTitle, IceCreamTotal)
	fmt.Printf("%s: $%d\n", MilkChocolateTitle, MilkChocolateTotal)
	fmt.Printf("%s: $%d\n", DoughnutTitle, DoughnutTotal)
	fmt.Printf("%s: $%d\n", PancakeTitle, PancakeTotal)
	fmt.Println("")
	var income int = BubblegumTotal + ToffeeTotal + IceCreamTotal + MilkChocolateTotal + DoughnutTotal + PancakeTotal
	fmt.Printf("Income: $%d\n", income)

	fmt.Println("Staff expenses:")
	var staffExpenses int
	fmt.Scanln(&staffExpenses)
	fmt.Println("Other expenses:")
	var otherExpenses int
	fmt.Scanln(&otherExpenses)
	fmt.Printf("Net income: $%d\n", income-staffExpenses-otherExpenses)
}
