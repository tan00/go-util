package util

import (

	//"github.com/json-iterator/go"

	//"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

// JSON writes json to response body.
func JSONMarshal(data interface{}, hasIndent bool) (content []byte, err error) {
	//var content []byte

	//use json lib github.com/json-iterator/go
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		return nil, err
	}
	return content, nil
}

// JSON writes json to response body.
func JSONUnmarshal(blob []byte, v interface{}) (err error) {
	//var content []byte
	//use json lib github.com/json-iterator/go
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(blob, v)
	return err
}
