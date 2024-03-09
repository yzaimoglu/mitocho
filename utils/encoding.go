package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"log"
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

func ToBase64(data interface{}) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return "", err
	}
	returnString := base64.StdEncoding.EncodeToString(buf.Bytes())
	return returnString, nil
}

func FromBase64(str string, target interface{}) error {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("Decode error: %v", err)
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(target)
}
