package config

import (
	"{{GoModule}}/internal/store/db"
	"{{GoModule}}/pkg/mysqls"
)

type Databases struct {
	Core Database `yaml:"core"`
}

type Database struct {
	Database string   `yaml:"database"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Port     string   `yaml:"port"`
	Masters  []string `yaml:"masters"`
	Slaves   []string `yaml:"slaves"`
}

func (c Databases) To() db.Config {
	return db.Config{
		Core: mysqls.Config{
			Database: c.Core.Database,
			User:     c.Core.User,
			Password: c.Core.Password,
			Port:     c.Core.Port,
			Masters:  c.Core.Masters,
			Slaves:   c.Core.Slaves,
		},
	}
}
