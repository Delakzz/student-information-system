package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadString(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func ReadInt(prompt string) int {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		return value
	}
}

func CheckInvalidID(id int) bool {
	return id > 0
}
