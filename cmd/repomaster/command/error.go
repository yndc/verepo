package command

import "fmt"

func PrintError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}
