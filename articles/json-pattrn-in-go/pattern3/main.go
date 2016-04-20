package main

import (
	"fmt"
)

// model
type User struct {
	Name          *string `json:"name"              validate:"nonzero,min=1"` // required, but no defaults
	Age           *uint   `json:"age,omitempty"     validate:"min=1"`         // optional
	Address       *string `json:"address,omitempty" validate:"min=1"`         // optional
	FavoriteColor string  `json:"favoriteColor"`                              // required, uses defaults
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
