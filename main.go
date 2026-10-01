package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, t - add tip, s - save bill, q - quit): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, _ := strconv.ParseFloat(price, 64)
		b.addItem(name, p)
		fmt.Println("Item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, _ := strconv.ParseFloat(tip, 64)
		b.updateTip(t)
		fmt.Println("Tip added - $", t)
		promptOptions(b)
	case "s":
		fmt.Println("Saving bill...")
		billFileName := b.name + ".txt"
		f, err := os.Create(billFileName)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		f.WriteString(b.format())
		fmt.Println("Bill saved to file - ", billFileName)
	case "q":
		fmt.Println("You have quit the app. Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please try again.")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
