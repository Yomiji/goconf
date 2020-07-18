package goconf_tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yomiji/goconf/v2"
	"github.com/yomiji/slog"
)

type Database struct {
	Server string
	Ports []int
	ConnectionMax int `toml:"connection_max"`
	Enabled bool
}
type TomlConfig struct {
	Database Database
}
type EnvConfig struct {
	Path string
	GoconfNum int `env:"GO_CONF_NUMBER"`
	GoconfFloat32 float32 `env:"GO_CONF_FLOAT32"`
	GoconfFloat64 float64 `env:"GO_CONF_FLOAT64"`
}

func TestMain(m *testing.M) {

	// set env variables
	if _, ok := os.LookupEnv("Path"); !ok {
		err:=os.Setenv("Path", "some/path")
		if err != nil {
			slog.Info("Failed to set Path")
		}
	}
	if err:=os.Setenv("GO_CONF_NUMBER", "123"); err != nil {
		slog.Info("Failed to set GO_CONF_NUMBER")
	}
	if err := os.Setenv("GO_CONF_FLOAT32", "12.3"); err != nil {
		slog.Info("Failed to set GO_CONF_FLOAT32")
	}
	if err := os.Setenv("GO_CONF_FLOAT64", "12.3"); err != nil {
		slog.Info("Failed to set GO_CONF_FLOAT64")
	}

	slog.ToggleLogging(true, true, true, true)

	os.Exit(m.Run())
}

func TestFileToml(t *testing.T) {
	conf := TomlConfig{}

	err := goconf.FromToml("test.toml", &conf)
	if err != nil {
		t.Fatalf("Failed to load .toml file: %v\n", err)
	}
	if conf.Database.Server == "" || len(conf.Database.Ports) <= 0 || conf.Database.ConnectionMax == 0  ||
		!conf.Database.Enabled {
		t.Fatal("Failed due to data not read from file")
	}
}

func TestByteToml(t *testing.T) {
	conf := TomlConfig{}
	f,err := os.Open("test.toml")
	if err != nil {
		t.Fatalf("Failed to load .toml file: %v\n", err)
	}
	all, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Failed to read .toml file: %v\n", err)
	}
	err = goconf.FromTomlBytes(all, &conf)
	if err != nil {
		t.Fatalf("Failed to translate .toml file: %v\n", err)
	}
	if conf.Database.Server == "" || len(conf.Database.Ports) <= 0 || conf.Database.ConnectionMax == 0  ||
		!conf.Database.Enabled {
		t.Fatalf("Failed due to data not read from file: %v", err)
	}
}
func TestReaderToml(t *testing.T) {
	conf := TomlConfig{}
	f,err := os.Open("test.toml")
	if err != nil {
		t.Fatalf("Failed to load .toml file: %v\n", err)
	}
	err = goconf.FromTomlReader(f, &conf)
	if err != nil {
		t.Fatalf("Failed to translate .toml file: %v\n", err)
	}
	if conf.Database.Server == "" || len(conf.Database.Ports) <= 0 || conf.Database.ConnectionMax == 0  ||
		!conf.Database.Enabled {
		t.Fatalf("Failed due to data not read from file: %v", err)
	}
}

func TestEnv(t *testing.T) {
	conf := EnvConfig{}

	err := goconf.FromEnvironment(&conf)
	if err != nil {
		t.Fatalf("Failed to load from environment: %v\n", err)
	}
	if conf.Path == "" || conf.GoconfFloat32 == 0 || conf.GoconfFloat64 == 0 || conf.GoconfNum == 0 {
		t.Fatal("Failed due to data not read from environment variables")
	}
}