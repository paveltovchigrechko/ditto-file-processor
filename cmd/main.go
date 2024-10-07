package main

import (
	"log"

	"github.com/paveltovchigrechko/ditto-file-processor/internal/ditto"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/helpers"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/validators"
)

const (
	inputDir  string = "../input/"
	outputDir string = "../output/"
)

func main() {
	files := ditto.ReadDittoFiles(inputDir)
	err := validators.ValidateFiles(files, inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := file.Name()
		projectName, localeName := ditto.SplitProjectAndLocale(fn)

		err := validators.ValidateNames(projectName, localeName, fn)
		if err != nil {
			continue
		}

		dittoJson := ditto.ExtractDittoKeys(inputDir+fn, projectName)

		encodedDitto := ditto.EncodeDittoKeys(dittoJson)

		// Might not work for Github action
		helpers.CreateDir(outputDir)
		ditto.CreateAndWriteJson(outputDir+localeName, encodedDitto)
	}
}

// Tests:
// 1. Add tests (parsing filename, file, key in map)
//     * Output file ()
// 2. Split the main() -> separate functions
// 3. Make this program as a part of pipeline: accept a single file (or filename) and output the result in std.Out
// 4. Think about Unicode in encoded JSON
