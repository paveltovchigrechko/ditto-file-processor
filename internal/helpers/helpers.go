package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/paveltovchigrechko/ditto-file-processor/internal/validators"
)

func CreateDir(path string) {
	dirExists, _ := validators.DirExists(path)

	if !dirExists {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)
		}
	}
}

func ReadArgs() []string {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) != 1 {
		log.Fatalf("Wrong number of arguments: expected only one file name")
	}

	return args
}
