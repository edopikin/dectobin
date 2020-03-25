package main

import (
	"strings"
	"fmt"

)

func main() {

	fmt.Println("Convert From Decimal (Base 10) To Binary (Base 2)")

	var number int

	for {

		fmt.Println("Kindly Input your number")
		fmt.Scan(&number)

		fmt.Printf("number = %d, in binary format = %b\n", number, number)


		fmt.Println("Do you want to go again; Yes or No")
		var response string
		fmt.Scan(&response)

		response = strings.ToLower(response)

		if response == "yes" {
			fmt.Println("Let's see what you got")
			continue
		} else if response == "no" {
			fmt.Println("Nothing for you")
			break
		}

	}
}