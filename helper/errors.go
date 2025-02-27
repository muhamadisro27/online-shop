package helper

import (
	"fmt"
	"os"
)

func PanicIfError(message string, err error) {
	fmt.Printf("%v %v\n", message, err)
	os.Exit(0)
}
