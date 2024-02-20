package utils

import (
	"encoding/json"
)

func ToJson(source interface{}) []byte {
	marshal, err := json.Marshal(source)
	if err != nil {
		return nil
	}
	return marshal
}

func FromJson(source []byte) interface{} {
	var destination interface{}
	err := json.Unmarshal(source, &destination)
	if err != nil {
		return nil
	}
	return destination
}

func FromJsonToObject(source []byte, destination interface{}) error {
	destinationLocal := FromJson(source)
	return InterfaceToObject(destinationLocal, destination)
}

func FromJsonToObjectMultiple(source [][]byte, destination interface{}) error {
	var destinationLocal []interface{}
	for _, sourceLocal := range source {
		destinationLocal = append(destinationLocal, FromJson(sourceLocal))
	}
	return InterfaceToObject(destinationLocal, destination)
}

func InterfaceToObject(input interface{}, output interface{}) error {
	jsonObj, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonObj, output)
	if err != nil {
		return err
	}

	return nil
}
