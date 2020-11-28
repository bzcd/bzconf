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
