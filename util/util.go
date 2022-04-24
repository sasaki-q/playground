package util

import (
	"fmt"
)

func Debug(arg any) {
	fmt.Print("DEBUG ==== ", arg)
}

func PrintError(err error) {
	fmt.Print("DEBUG error message === ", err)
}
