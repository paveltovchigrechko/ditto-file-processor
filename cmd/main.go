package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/paveltovchigrechko/ditto-file-processor/internal/ditto"
)

const (
	inputDir  string = "./input/"
	outputDir string = "./output/"

	prefix string = ""
	indent string = "    "
)

func main() {
	// Catch the project name (it is the key in the resulting map) and the output file name (locale)
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		projectName, localeName := ditto.SplitProjectAndLocale(file.Name())

		// Open file
		f, err := os.ReadFile(inputDir + file.Name())
		if err != nil {
			log.Println(err)
		}

		// Decode file and save it as map[projectName]json_blob
		var m map[string]interface{}
		err = json.Unmarshal(f, &m)
		if err != nil {
			log.Println(err)
		}

		// Extract JSON blob
		jsonBlob := m[projectName] // Handle error

		// Encode the blob into JSON
		encodedBlob, _ := json.MarshalIndent(jsonBlob, prefix, indent)
		// fmt.Print(string(encodedBlob))
		if err != nil {
			log.Println(err)
		}

		// Create a JSON file and save JSON blob into it
		// Might not work for Github action
		err = os.Mkdir(outputDir, 0777)
		if err != nil {
			log.Println(err)
		}

		newFile, err := os.Create(outputDir + localeName)
		// Why it is output in console
		if err != nil {
			log.Println(err)
		}

		_, err = newFile.Write(encodedBlob)
		if err != nil {
			log.Println(err)
		}
	}
}

// Tests:
// 1. Add tests (parsing filename, file, key in map)
//     * Output file ()
// 2. Split the main() -> separate functions
// 3. Make this program as a part of pipeline: accept a single file (or filename) and output the result in std.Out
// 4. Think about Unicode in encoded JSON
