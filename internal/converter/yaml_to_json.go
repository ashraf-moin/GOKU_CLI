package converter

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func YAMLToJSON(data []byte) ([]byte, error) {
	var obj interface{}

	err := yaml.Unmarshal(data, &obj)

	if err != nil {
		return nil, err
	}
	//MarshalIndent and Marshal MarshalIndent(Go object into pretty JSON format).
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
