package logging

import (
	"fmt"
	"strings"
)

func PanicStack(err error) {
	trErr := UnwrapStacktrace(err)
	if trErr == nil {
		panic(err)
	}

	st := trErr.StackTrace()
	if len(st) == 0 {
		panic(err)
	}

	message := fmt.Sprintf("%s %+v", err.Error(), st[0:1])
	message = strings.ReplaceAll(message, "\n", " ")
	message = strings.ReplaceAll(message, "\t", "")

	panic(message)
}
