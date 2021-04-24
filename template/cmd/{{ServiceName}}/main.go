package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	ic "{{GoModule}}/internal/config"
	"{{GoModule}}/internal/{{ServiceName}}"
	"{{GoModule}}/internal/rest"
	"{{GoModule}}/internal/store/db"
	"{{GoModule}}/pkg/config"
	"{{GoModule}}/pkg/server"
)

const configPathEnvKey = "CONFIG_PATH"

type conf struct {
	ListenAddr      string `yaml:"listen_addr"`
	DebugListenAddr string `yaml:"debug_listen_addr"`
	Swagger         struct {
		Host string `yaml:"host"`
	} `yaml:"swagger"`
	Debug bool         `yaml:"debug"`
	DB    ic.Databases `yaml:"databases"`
}

func (c conf) isValid() error {
	return nil
}

func main() {
	flag.Parse()

	c, err := parseConfig()
	if err != nil {
		panic(err)
	}

	initGlobalLog(c.Debug)

	dbs, err := db.New(c.DB.To())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init databases")
	}

	{{ServiceName}}s := {{ServiceName}}.New(dbs)

	ctx := context.Background()
	go func() {
		err = server.GracefullyListenAndServe(ctx, c.DebugListenAddr, rest.RegisterDebugRouter(c.Swagger.Host))
		log.Err(err).Msg("debug server exists")
	}()
	handler := rest.RegisterRouter({{ServiceName}}s)
	err = server.GracefullyListenAndServe(ctx, c.ListenAddr, handler)
	log.Err(err).Msg("app exits")
}

func initGlobalLog(debug bool) {
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func parseConfig() (conf, error) {
	p := strings.TrimSpace(os.Getenv(configPathEnvKey))
	if p == "" {
		return conf{}, fmt.Errorf("%s should be set to config file path", configPathEnvKey)
	}
	var c conf
	if err := config.ParseYAML(p, &c); err != nil {
		return conf{}, err
	}
	if err := c.isValid(); err != nil {
		return conf{}, err
	}
	return c, nil
}
