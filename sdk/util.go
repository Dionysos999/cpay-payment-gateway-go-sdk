package sdk

import (
	"github.com/mitchellh/mapstructure"
)

func convStruct2Map(p interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := mapstructure.Decode(p, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
