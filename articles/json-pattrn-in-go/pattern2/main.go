package main

import (
	"fmt"
)

// model
	type User struct {
		Name          *string `json:"name"`              // required, but no defaults
		Age           *uint   `json:"age,omitempty"`     // optional
		Address       *string `json:"address,omitempty"` // optional
		FavoriteColor string  `json:"favoriteColor"`     // required, uses defaults
	}

func main(){
	// unmarshalling
	var user User

	if err := json.NewDecoder(jsonByteSlice).Decode(&user); err != nil {

	}

	// marshalling
	if jsonByteSlice, err := json.Marshal(object); err != nil {

	}

	// validation
	func Validate(user User) {
		// default - validate value
		// optional - if non nil, validate value
		// required non default - validate not nil, then validate value
	}
}

