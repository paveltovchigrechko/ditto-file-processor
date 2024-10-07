package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paveltovchigrechko/ditto-file-processor/internal/ditto"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/helpers"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/validators"
)

func main() {
	file := helpers.ReadArgs()[0]
	err := validators.ValidateFile(file)
	if err != nil {
		log.Fatal(err)
	}

	projectName, localeName := ditto.SplitProjectAndLocale(file)

	err = validators.ValidateNames(projectName, localeName, file)
	if err != nil {
		log.Fatal(err)
	}

	dittoJson := ditto.ExtractDittoKeys(file, projectName)
	converted := []byte(fmt.Sprintf("%v", dittoJson))
	os.Stdout.Write(converted)
}

// Tests:
// 1. Add tests (parsing filename, file, key in map)
//     * Output file ()
// 2. Split the main() -> separate functions
// 3. Make this program as a part of pipeline: accept a single file (or filename) and output the result in std.Out
// 4. Think about Unicode in encoded JSON
