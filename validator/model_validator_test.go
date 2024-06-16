package validator

import (
	"testing"
)

func Test__ValidateModel__EmptyFile(t *testing.T) {
	_, fileError := ValidateModel("sample.empty")
	if fileError == nil {
		t.Error(
			"ValidateModel 함수는 텅빈 파일을 읽을 때, 에러를 반환해야 합니다.",
		)
	}
}

func Test__ValidateModel__InvalidJsonFile(t *testing.T) {
	_, fileError := ValidateModel("sample.invalid.json")
	if fileError == nil {
		t.Error(
			"ValidateModel 함수는 invalid json 파일을 읽을 때, 에러를 반환해야 합니다.",
		)
	}
}


func Test__ValidateModel__ValidJsonFile(t *testing.T) {
	_, fileError := ValidateModel("sample.json")
	if fileError != nil {
		t.Error(
			"ValidateModel 함수는 valid json 파일을 읽을 때, 에러를 반환하지 않아야 합니다.",
		)
	}
}

func Test__ValidateModel__ValidYamlFile(t *testing.T) {
	_, fileError := ValidateModel("sample.yaml")
	if fileError != nil {
		t.Error(
			"ValidateModel 함수는 valid yaml 파일을 읽을 때, 에러를 반환하지 않아야 합니다.",
		)
	}
}