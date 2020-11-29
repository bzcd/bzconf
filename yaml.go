package bzconf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadYaml load yaml file, and decode to out
func LoadYaml(file string, out interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

// LoadYamlPath load yaml file, and decode to out
func LoadYamlPath(path string, out interface{}, opts ...*OptionsList) error {
	return loadPath(path, out, LoadYaml, opts...)
}
