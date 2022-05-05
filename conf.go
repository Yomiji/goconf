package goconf

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func unmarshal(b []byte, obj interface{}) error {
	err := toml.Unmarshal(b, obj)
	if err != nil {
		return err
	}

	return nil
}

func FromToml(fileName string, obj interface{}) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return unmarshal(b, obj)
}

func FromTomlReader(reader io.Reader, obj interface{}) error {
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	return unmarshal(b, obj)
}

func FromTomlBytes(b []byte, obj interface{}) error {
	return unmarshal(b, obj)
}
