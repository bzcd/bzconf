package bzconf

import (
	"io/ioutil"
	"strings"
)

// ListAll list all file at the path
func ListAll(path string, opts ...*OptionsList) ([]string, error) {
	var opt *OptionsList
	if len(opts) == 0 {
		opt = DefaultOptionsList()
	} else {
		opt = opts[0]
	}

	if len(path) == 0 {
		path = "./"
	} else if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	return listAll(path, opt, 1)
}

// listAll list all file at the path
func listAll(path string, opt *OptionsList, depth int) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, f := range files {
		name := f.Name()
		fa := path + f.Name()

		if f.IsDir() {
			if opt.IncludeDir {
				res = append(res, fa)
			}

			if opt.Recursion {
				if opt.Depth == 0 || opt.Depth > depth {
					r, err := listAll(fa+"/", opt, depth+1)
					if err != nil {
						if opt.IgnoreError {
							continue
						}

						return nil, err
					}

					if len(r) > 0 {
						res = append(res, r...)
					}
				}
			}
		} else if f.Mode().IsRegular() {
			if len(opt.Prefix) != 0 && !strings.HasPrefix(name, opt.Prefix) {
				continue
			}

			if len(opt.Suffix) != 0 && !strings.HasSuffix(name, opt.Suffix) {
				continue
			}

			res = append(res, fa)
		}

	}

	return res, nil
}

// loadPath load files, and decode to out
func loadPath(path string, out interface{}, loadFunc func(string, interface{}) error, opts ...*OptionsList) error {
	l, err := ListAll(path, opts...)
	if err != nil {
		return err
	}

	for _, f := range l {
		if err = loadFunc(f, out); err != nil {
			return err
		}
	}

	return nil
}
