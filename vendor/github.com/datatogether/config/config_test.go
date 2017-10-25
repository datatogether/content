package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	path, err := writeTempFile(".env.test", `
PORT=3000
BOOL=true
FALSE_BOOL=false
TEST_STRING=test_string
TEST_STRING_SLICE=test,string,slice
`)

	if err != nil {
		t.Errorf("error writing file: %s", err)
		return
	}

	cfg := struct {
		Port            int
		Bool            bool
		FalseBool       bool
		TestString      string
		TestStringSlice []string
	}{}

	if err := Load(&cfg, path); err != nil {
		t.Errorf("error loading file: %s", err)
		return
	}

	if cfg.Port != 3000 {
		t.Errorf("Port value mismatch. %d != %d", cfg.Port, 3000)
	}

	if cfg.TestString != "test_string" {
		t.Errorf("TestString value mismatch. %s != %s", cfg.TestString, "test_string")
	}
	if !cfg.Bool {
		t.Errorf("Bool value mismatch. %t != %t", cfg.Bool, true)
	}
	if cfg.FalseBool {
		t.Errorf("FalseBool value mismatch. %t != %t", cfg.FalseBool, false)
	}

	if len(cfg.TestStringSlice) != 3 {
		t.Errorf("TestString value mismatch. %v != %v", cfg.TestStringSlice, []string{"test", "string", "slice"})
	}
}

func TestEnvVarKey(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Config", "CONFIG"},
		{"ConfigVar", "CONFIG_VAR"},
		{"configVar", "CONFIG_VAR"},
		{"configvar", "CONFIGVAR"},
		{"AckConfigVar", "ACK_CONFIG_VAR"},
	}

	for i, c := range cases {
		if got := EnvVarKey(c.in); got != c.out {
			t.Errorf("case %d mismatch: %s != %s", i, c.out, got)
		}
	}
}

func writeTempFile(filename, data string) (string, error) {
	path := filepath.Join(os.TempDir(), filename)
	err := ioutil.WriteFile(path, []byte(data), os.ModePerm)
	return path, err
}
