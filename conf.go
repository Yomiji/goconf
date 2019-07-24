package goconf

import (
	"github.com/pelletier/go-toml"
)


type ConfigFile struct {}

func (cfg ConfigFile) Load(fileName string, obj interface{}) error {
	t,err := toml.LoadFile(fileName)
	if err != nil {
		return err
	}
	return t.Unmarshal(obj)
}