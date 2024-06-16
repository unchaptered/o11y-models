package validator

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2" // YAML 패키지 임포트
)

// func GetFileFormat(filename string) (string, error) {
func ValidateModel(filename string) (interface{}, error) {
	var fileData []byte
	var fileError error
	
	fileData, fileError = os.ReadFile(filename)
	if fileError != nil {
		return nil, fileError
	}

	// var jsonData interface{}
	var jsonData interface{}
	var yamlData interface{}
	jsonError := json.Unmarshal(fileData, &jsonData)
	yamlError := yaml.Unmarshal(fileData, &yamlData)
	if jsonError != nil && yamlError != nil {
		return nil, jsonError
	}
	
	if jsonError != nil {
		return yamlData, nil
	}

	return jsonData, nil
}