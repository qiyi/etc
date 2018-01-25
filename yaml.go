package etc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func NewYamlSource(name string, file string) (*MapSource, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	y := make(map[string]interface{})
	err = yaml.Unmarshal(bytes, &y)
	if err != nil {
		return nil, err
	}
	return NewMapSource(name, y), nil
}
