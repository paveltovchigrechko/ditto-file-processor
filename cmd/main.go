package main

import (
	"log"

	"github.com/paveltovchigrechko/ditto-file-processor/internal/ditto"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/helpers"
	"github.com/paveltovchigrechko/ditto-file-processor/internal/validators"
)

const (
	inputDir  string = "input/"
	outputDir string = "output/"
)

func main() {
	dh := ditto.New(ditto.OSWrapper{})
	files, err := dh.ReadDittoFiles(inputDir)
	if err != nil {
		log.Fatal(err)
	}
	err = validators.ValidateFiles(files, inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := file.Name()
		projectName, localeName, err := ditto.SplitProjectAndLocale(fn)
		if err != nil {
			log.Println(err)
		}

		err = validators.ValidateNames(projectName, localeName, fn)
		if err != nil {
			log.Println(err)
			continue
		}

		dittoJson, err := ditto.ExtractDittoKeys(inputDir+fn, projectName)
		if err != nil {
			log.Println(err)
			continue
		}

		encodedDitto, err := ditto.EncodeDittoKeys(dittoJson)
		if err != nil {
			log.Println(err)
			continue
		}

		err = helpers.CreateDir(outputDir)
		if err != nil {
			log.Println(err)
			continue
		}

		err = ditto.CreateAndWriteJson(outputDir+localeName, encodedDitto)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
