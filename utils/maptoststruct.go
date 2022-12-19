package utils

import "encoding/json"

func MapToStruct(in interface{}, out interface{}) error {
	marshal, err := json.Marshal(in)
	if err != nil {
		return err
	}
	json.Unmarshal(marshal, out)
	return nil
}
