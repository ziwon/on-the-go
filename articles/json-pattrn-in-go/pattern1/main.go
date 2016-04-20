package main

import (
	"encode/json"
	"fmt"
)

// model
type User struct {
	Name    string `json:"name"    validate:"nonzero"`
	Age     uint   `json:"age"     validate:"min=1"`
	Address string `json:"address" validate:"nonzero"`
}

func main() {
	// unmarshalling
	var user User
	if err := json.NewDecoder(jsonByteSlice).Decode(&user); err != nil {

	}

	// marshalling
	if jsonByteSlice, err := json.Marshal(object); err != nil {

	}

	// validation
	if errs := validator.Validate(user); errs != nil {

	}
}
