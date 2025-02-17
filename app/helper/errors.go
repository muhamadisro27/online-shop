package helper

import (
	"fmt"
	"os"
)

func PanicIfError(err error) {
	fmt.Println(err)
	os.Exit(0)
}
