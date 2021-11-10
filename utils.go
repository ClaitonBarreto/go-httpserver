package main

import (
	"encoding/json"
	"io"
)

func JsonDecoder(jsonObject io.ReadCloser) (User, error) {
	decoder := json.NewDecoder(jsonObject)
	var user User

	err := decoder.Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
