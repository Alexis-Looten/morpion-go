package morpion

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadNumber() (number string, err error) {
	valid := false
	for !valid {
		fmt.Println("Please, select a case by typing his number.")
		fmt.Print("What is your choice ? ")
		number, err = reader.ReadString('\n')
		if err != nil {
			return number, err
		}
		number = strings.TrimSpace(number)
		if len(number) != 1 {
			fmt.Printf("Invalid number input.  number =%v, len=%v\n", number, len(number))
			continue
		}

		figures := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

		for i := range figures {
			if number != figures[i] {
				fmt.Printf("Invalid number input.  number =%v, len=%v\n", number, len(number))
				continue
			}
		}

		valid = true
	}
	return number, nil
}
