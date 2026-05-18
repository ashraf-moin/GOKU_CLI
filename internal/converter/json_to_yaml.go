package converter

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func JSONToYAML(data []byte) ([]byte, error) {
	var obj interface{}
//convert data into object
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}