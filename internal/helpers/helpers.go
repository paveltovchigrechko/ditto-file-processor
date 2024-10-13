package helpers

import (
	"fmt"
	"log"
	"os"
)

func ReadArgs() []string {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) != 1 {
		log.Fatalf("Wrong number of arguments: expected only one file name")
	}

	return args
}
