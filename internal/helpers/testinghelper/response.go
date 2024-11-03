package testinghelper

import (
	"testing"

	"github.com/buger/jsonparser"
)

type (
	PropertyTest struct {
		Type  jsonparser.ValueType
		Value any
	}

	DataTests map[string]PropertyTest
)

func CheckHttpStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Incorrect HTTP Status, got: %d, want: %d", got, want)
	}
}

func CheckSuccess(t *testing.T, resBody []byte) {
	isSuccess, err := jsonparser.GetBoolean(resBody, "is_success")

	if err != nil {
		t.Error("is_success doesn't exist")
		return
	}

	if !isSuccess {
		t.Error("It should get a success response, but it doesn't.")
	}
}

func CheckData(t *testing.T, resBody []byte, tests DataTests, keys ...string) {
	if len(keys) == 0 {
		keys = append(keys, "data")
	}

	err := jsonparser.ObjectEach(resBody, func(k, v []byte, vt jsonparser.ValueType, o int) error {
		key := string(k)
		val := string(v)

		if test, ok := tests[key]; ok {
			if vt != test.Type {
				t.Errorf("Incorrect type of %s, got: %v, want: %v", key, vt, test.Type)
			}
			if test.Value != nil && val != test.Value {
				t.Errorf("Incorrect value of %s, got: %v, want: %v", key, val, test.Value)
			}
		}

		return nil
	}, keys...)

	if err != nil {
		t.Error("data doesn't exist")
	}
}
