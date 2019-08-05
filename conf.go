package goconf

import (
	"github.com/pelletier/go-toml"
	"io"
)

func FromToml(fileName string, obj interface{}) error {
	t, err := toml.LoadFile(fileName)
	if err != nil {
		return err
	}
	return t.Unmarshal(obj)
}

func FromTomlReader(reader io.Reader, obj interface{}) error {
	t, err := toml.LoadReader(reader)
	if err != nil {
		return err
	}
	return t.Unmarshal(obj)
}

func FromTomlBytes(bytes []byte, obj interface{}) error {
	t, err := toml.LoadBytes(bytes)
	if err != nil {
		return err
	}
	return t.Unmarshal(obj)
}
