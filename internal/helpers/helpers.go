package helpers

import (
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
