package bzconf

import "testing"

func TestLoadYaml(t *testing.T) {
	var v struct {
		App     string `yaml:"app"`
		Version string `yaml:"version"`
	}

	LoadYaml("testdata/app.yaml", &v)
	t.Logf("%+v\n", v)
}

func TestLoadYamlPath(t *testing.T) {
	var v struct {
		App     string `yaml:"app"`
		Version string `yaml:"version"`

		Info struct {
			Desc string                 `yaml:"desc"`
			Body map[string]interface{} `yaml:"body"`
		} `yaml:"info"`
	}

	err := LoadYamlPath("testdata", &v)
	t.Logf("err[%v], %+v\n", err, v)
}
