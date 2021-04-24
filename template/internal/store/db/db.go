package db

import (
	"github.com/linxGnu/mssqlx"

	"{{GoModule}}/pkg/mysqls"
)

type Manager struct {
	core *mssqlx.DBs
}

type Config struct {
	Core mysqls.Config
}

func New(c Config, opts ...mysqls.Option) (*Manager, error) {
	var err error
	var m Manager

	m.core, err = mysqls.New(c.Core, opts...)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
