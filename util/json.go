package util

import (
	"encoding/json"
)

func JsonBack(s []byte) ([]byte, error) {
	temp := new(interface{})

	err := json.Unmarshal(s, temp)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(temp)
	return result, err
}
