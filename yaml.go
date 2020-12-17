package bzconf

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadYaml load yaml file, and decode to out
func LoadYaml(file string, out interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return LoadYamlBuffer(data, out)
}

// LoadYamlBuffer load yaml buffer, and decode to out
func LoadYamlBuffer(buffer []byte, out interface{}) error {
	if len(buffer) == 0 {
		return errors.New("empty buffer")
	}

	return yaml.Unmarshal(buffer, out)
}

// LoadYamlPath load yaml file, and decode to out
func LoadYamlPath(path string, out interface{}, opts ...*OptionsList) error {
	opts = append(opts, &OptionsList{IncludeDir: false})
	return loadPath(path, out, LoadYaml, opts...)
}
