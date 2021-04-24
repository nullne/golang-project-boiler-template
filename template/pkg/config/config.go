package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseYAML(path string, obj interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	return decoder.Decode(obj)
}

func WriteYAML(w io.Writer, obj interface{}) error {
	bs, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(bs)
	return err
}
