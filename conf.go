package goconf

import (
	"github.com/pelletier/go-toml"
)

func FromToml(fileName string, obj interface{}) error {
	t,err := toml.LoadFile(fileName)
	if err != nil {
		return err
	}
	return t.Unmarshal(obj)
}