package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"perso.go/morpion/morpion"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func RetryOrExit() (err error) {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print("Do you want to replay ? ")

	retrychoice, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	retrychoice = strings.TrimSpace(retrychoice)
	switch retrychoice {
	case "yes":
		CallClear()
		main()
	case "no":
		os.Exit(0)
	default:
		fmt.Printf("Your answer is't correct, please retry with the words 'yes' or 'no' : %v\n", retrychoice)
		RetryOrExit()
	}
	return nil
}

func main() {

	g, _ := morpion.New()

	choice := ""
	for {
		morpion.Draw(g, choice)

		switch g.State {
		case "End", "PlayerWin":
			err2 := RetryOrExit()
			if err2 != nil {
				fmt.Printf("Could not understand your choice: %v\n", err2)
				os.Exit(1)
			}
		}

		l, err := morpion.ReadNumber()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v\n", err)
			os.Exit(1)
		}
		choice = l
		g.MakeAChoice(choice)
		CallClear()
	}

}
