package main

import (
	"fmt"

	"o11y-models/validator"

	"github.com/xeipuuv/gojsonschema"
)

func main() {

	fileNmae := "file:///Users/leeminseok/Downloads/dev-files/o11y-tool/o11y-models/validator/sample.json"
	validator.ValidateModel(fileNmae)

	schemaFile := "file:///Users/leeminseok/Downloads/dev-files/o11y-tool/o11y-models/validator/sample.json"
    schemaLoader := gojsonschema.NewReferenceLoader(schemaFile)
	fmt.Print(schemaLoader.LoadJSON())

}